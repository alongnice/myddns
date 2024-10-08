package config

import (
	"fmt"
	"log"
	"myddns/util"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Webhook struct {
	WebhookURL         string
	WebhookRequestBody string
}

// updateStatusType 更新状态
type updateStatusType string

const (
	// UpdatedNothing 未改变
	UpdatedNoChange updateStatusType = "未改变"
	// UpdatedFailed 更新失败
	UpdatedFail = "失败"
	// UpdatedSuccess 更新成功
	UpdatedSuccess = "成功"
)

func getDomainsStatus(domains []*Domain) updateStatusType {
	successNum := 0
	for _, v46 := range domains {
		switch v46.UpdateStatus {
		case UpdatedFail:
			return UpdatedFail
			// 如果有一个失败 则结果为失败
		case UpdatedSuccess:
			successNum++
		}
	}
	if successNum > 0 {
		return UpdatedSuccess
	}
	return UpdatedNoChange
}

// 逗号分隔
func getDomainsStr(domains []*Domain) string {
	str := ""
	for i, v46 := range domains {
		str += v46.String()
		if i != len(domains)-1 {
			str += ", "
		}
	}
	return str
}

// 替换参数
func (domains *Domains) replacePara(orgPara string, ipv4Result updateStatusType, ipv6Result updateStatusType) (newPara string) {
	orgPara = strings.ReplaceAll(orgPara, "#{ipv4Addr}", domains.Ipv4Addr)
	orgPara = strings.ReplaceAll(orgPara, "#{ipv4Result}", string(ipv4Result))
	orgPara = strings.ReplaceAll(orgPara, "#{ipv4Domains}", getDomainsStr(domains.Ipv4Domains))

	orgPara = strings.ReplaceAll(orgPara, "#{ipv6Addr}", domains.Ipv6Addr)
	orgPara = strings.ReplaceAll(orgPara, "#{ipv6Result}", string(ipv6Result))
	orgPara = strings.ReplaceAll(orgPara, "#{ipv6Domains}", getDomainsStr(domains.Ipv6Domains))

	return orgPara
}

// 添加或者更新 IPv4/IPv6 域名
func ExecWebhook(domains *Domains, conf *Config) {
	v4Status := getDomainsStatus(domains.Ipv4Domains)
	v6Status := getDomainsStatus(domains.Ipv6Domains)

	if conf.WebhookURL != "" && (v4Status == UpdatedNoChange || v6Status == UpdatedNoChange) {
		// 成功或者失败都触发webhook
		method := "GET"
		postPara := ""
		contentType := "application/x-www-form-urlencoded"
		if conf.WebhookRequestBody != "" {
			method = "POST"
			postPara = domains.replacePara(conf.WebhookRequestBody, v4Status, v6Status)
			contentType = "application/json"
		}
		requestURL := domains.replacePara(conf.WebhookURL, v4Status, v6Status)
		u, err := url.Parse(requestURL)
		if err != nil {
			log.Println("webhook中的URL不正确")
			return
		}
		req, err := http.NewRequest(method, fmt.Sprintf("%s://%s%s?%s", u.Scheme, u.Host, u.Path, u.Query().Encode()), strings.NewReader(postPara))
		if err != nil {
			log.Println("创建webhook请求失败", err)
			return
		}

		req.Header.Add("content-type", contentType)

		clt := http.Client{}
		clt.Timeout = 30 * time.Second
		resp, err := clt.Do(req)
		body, err := util.GetHTTPResponseOrg(resp, requestURL, err)
		if err == nil {
			log.Println(fmt.Sprintf("webhook 被调用成功,返回数据: %s", string(body)))
		} else {
			log.Println("webhook 调用失败", err)
		}
	}
}
