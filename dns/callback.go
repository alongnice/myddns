package dns

import (
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Callback struct {
	DNSConfig config.DNSConfig
	Domains   config.Domains
	TTL       string
}

func (cb *Callback) Init(conf *config.Config) {
	cb.DNSConfig = conf.DNS
	cb.Domains.ParseDomain(conf)
	if conf.TTL == "" {
		cb.TTL = "600"
	} else {
		cb.TTL = conf.TTL
	}
}

func (cb *Callback) AddUpdateDomainRecords() config.Domains {
	cb.addUpdateDomainRecords("A")
	cb.addUpdateDomainRecords("AAAA")
	return cb.Domains
}

var lastIpv4 string
var lastIpv6 string

func replacePara(orgPara, ipAddr string, domain *config.Domain, recordType string, ttl string) (newPara string) {
	orgPara = strings.ReplaceAll(orgPara, "#{ip}", ipAddr)
	orgPara = strings.ReplaceAll(orgPara, "#{domain}", domain.DomainName)
	orgPara = strings.ReplaceAll(orgPara, "#{subdomain}", recordType)
	orgPara = strings.ReplaceAll(orgPara, "#{ttl}", ttl)
	return orgPara
}

func (cb *Callback) addUpdateDomainRecords(recordType string) {
	ipAddr, domains := cb.Domains.ParseDomainResult(recordType)

	if ipAddr == "" {
		return
	}
	if recordType == "A" {
		if ipAddr == lastIpv4 {
			log.Println("你的ipv4地址未发生变化，没有触发callback")
			return
		}
		lastIpv4 = ipAddr
	} else {
		if ipAddr == lastIpv6 {
			log.Println("你的ipv6地址未发生变化，没有触发callback")
			return
		}
		lastIpv6 = ipAddr
	}

	for _, domain := range domains {
		method := "GET"
		postPara := ""
		contentType := "application/x-www-form-urlencoded"
		if cb.DNSConfig.Secret != "" {
			method = "POST"
			postPara = replacePara(cb.DNSConfig.ID, ipAddr, domain, recordType, cb.TTL)
			contentType = "application/x-www-form-urlencoded"
		}
		requestURL := replacePara(cb.DNSConfig.ID, ipAddr, domain, recordType, cb.TTL)
		u, err := url.Parse(requestURL)
		if err != nil {
			log.Panicln("callback 的URL格式错误")
			return
		}
		req, err := http.NewRequest(method, u.String(), strings.NewReader(postPara))
		if err != nil {
			log.Println("callback 请求异常")
			return
		}
		req.Header.Add("Content-Type", contentType)

		client := http.Client{}
		client.Timeout = 30 * time.Second
		resp, err := client.Do(req)
		body, err := util.GetHTTPResponseOrg(resp, u.String(), err)
		if err == nil {
			log.Println("callback 请求成功", string(body))
			domain.UpdateStatus = config.UpdatedSuccess
		} else {
			log.Println("callback 请求失败", string(body), "错误信息：", err)
			domain.UpdateStatus = config.UpdatedFail
		}
	}
}
