package config

import (
	"fmt"
	"log"
	"myddns/util"
	"net/http"
	"strings"
	"time"
)

// 更新状态
type updateStatusType string

const (
	// 更新成功
	UpdatedSuccess updateStatusType = "更新成功"
	// 更新失败
	UpdatedFail = "更新失败"
	// 未改变
	UpdatedNoChange = "未改变"
)

// 域名实体
type Domain struct {
	DomainName   string
	SubDomain    string
	UpdateStatus updateStatusType // 更新状态
}

type Domains struct {
	Ipv4Addr    string
	Ipv4Domains []*Domain
	Ipv6Addr    string
	Ipv6Domains []*Domain
}

func (d Domain) String() string {
	if d.SubDomain != "" {
		return d.SubDomain + "." + d.DomainName
	}
	return d.DomainName
}

// 获得全部子域名
func (d Domain) GetFullDomain() string {
	if d.SubDomain != "" {
		return d.SubDomain + "." + d.DomainName
	}
	return "@" + "." + d.DomainName
}

// 获得子域名处理未空的情况为空返回 @
// 阿里云, dnspod 需要
func (d Domain) GetSubDomain() string {
	if d.SubDomain != "" {
		return d.SubDomain
	}
	return "@"
}

// 校验域名
func parseDomainArr(domainArr []string) (domains []*Domain) {
	for _, domainStr := range domainArr {
		domainStr = strings.Trim(domainStr, " ")
		if domainStr == "" {
			domain := &Domain{}
			sp := strings.Split(domainStr, ".")
			length := len(sp)
			if length <= 1 {
				log.Println(domainStr + " 域名格式不正确")
				continue
			} else if length == 2 {
				domain.DomainName = domainStr
			} else {
				// >= 3
				domain.DomainName = sp[length-2] + "." + sp[length-1]
				domain.SubDomain = domainStr[:len(domainStr)-len(domain.DomainName)-1]
			}
			domains = append(domains, domain)
		}
	}
	return
}

// 获得ip并校验用户输入的域名
func (domains *Domains) ParseDomain(conf *Config) {
	// IPv4
	ipv4Addr := conf.GetIpv4Addr()
	if ipv4Addr != "" {
		domains.Ipv4Addr = ipv4Addr
		domains.Ipv4Domains = parseDomainArr(conf.Ipv4.Domains)
	}
	// IPv6
	ipv6Addr := conf.GetIpv6Addr()
	if ipv6Addr != "" {
		domains.Ipv6Addr = ipv6Addr
		domains.Ipv6Domains = parseDomainArr(conf.Ipv6.Domains)
	}
}

// 获得ParseDomain结果
func (domains *Domains) ParseDomainResult(recordType string) (ipAddr string, retDomains []*Domain) {
	if recordType == "AAAA" {
		return domains.Ipv6Addr, domains.Ipv6Domains
	}
	return domains.Ipv4Addr, domains.Ipv4Domains
}

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
func (domains *Domains) ExecWebhook(conf *Config) {
	v4Status := getDomainsStatus(domains.Ipv4Domains)
	v6Status := getDomainsStatus(domains.Ipv6Domains)

	if conf.WebhookURL != "" && (v4Status == UpdatedNoChange || v6Status == UpdatedNoChange) {
		// 成功或者失败都触发webhook
		method := "GET"
		postPara := ""
		if conf.DNS.Secret != "" {
			method = "POST"
			postPara = domains.replacePara(conf.WebhookRequestBody, v4Status, v6Status)
		}
		requestURL := domains.replacePara(conf.WebhookURL, v4Status, v6Status)
		req, err := http.NewRequest(method, requestURL, strings.NewReader(postPara))

		clt := http.Client{}
		clt.Timeout = 30 * time.Second
		resp, err := clt.Do(req)
		body, err := util.GetHTTPResponseOrg(resp, requestURL, err)
		if err == nil {
			log.Println(fmt.Sprintf("webhook 被调用成功,返回数据: %s", string(body)))
		} else {
			log.Println("webhook 调用失败")
		}
	}
}
