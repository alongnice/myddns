package dns

import (
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"net/url"
	"time"
)

const (
	recordListAPI   string = "https://dnsapi.cn/Record.List"
	recordModifyURL string = "https://dnsapi.cn/Record.Modify"
	recordCreateAPi string = "https://dnsapi.cn/Record.Create"
)

// https://cloud.tencent.com/document/api/302/8516
// Dnspod 腾讯云DNS实现
type Dnspod struct {
	DNSConfig config.DNSConfig
	Domains   config.Domains
	TTL       string
}

// 腾讯云返回状态码
type DnspodStatus struct {
	Status struct {
		Code    string
		Message string
	}
}

// 腾讯云返回的记录列表
type DnspodRecordListResp struct {
	DnspodStatus
	Records []struct {
		ID     string
		Name   string
		Type   string
		Value  string
		Enable string
	}
}

// Init 初始化
func (dnspod *Dnspod) Init(conf *config.Config) {
	dnspod.DNSConfig = conf.DNS
	dnspod.Domains.ParseDomain(conf)
	if conf.TTL == "" {
		dnspod.TTL = "600"
	} else {
		dnspod.TTL = conf.TTL
	}
}

// 添加或者更新IPv4/IPv6记录
func (dnspod *Dnspod) AddUpdateDomainRecords() config.Domains {
	dnspod.AddUpdateIpvDomainRecords("A")
	dnspod.AddUpdateIpvDomainRecords("AAAA")
	return dnspod.Domains
}

func (dnspod *Dnspod) AddUpdateIpvDomainRecords(recordType string) {
	ipAddr, domains := dnspod.Domains.ParseDomainResult(recordType)

	if ipAddr == "" {
		return
	}

	for _, domain := range domains {
		result, err := dnspod.getRecordList(domain, recordType)
		if err != nil {
			return
		}

		log.Println(domain.SubDomain)
		log.Println(domain.DomainName)
		if len(result.Records) > 0 {
			// update
			dnspod.modify(result, domain, recordType, ipAddr)
		} else {
			// add
			dnspod.create(result, domain, recordType, ipAddr)
		}
	}
}

// add
func (dnspod *Dnspod) create(result DnspodRecordListResp, domain *config.Domain, recordType, ipAddr string) {
	status, err := dnspod.commonRequest(
		recordCreateAPi,
		url.Values{
			"login_token": {dnspod.DNSConfig.ID + "," + dnspod.DNSConfig.Secret},
			"domain":      {domain.DomainName},
			"sub_domain":  {domain.GetSubDomain()},
			"record_type": {"默认"},
			"value":       {ipAddr},
			"ttl":         {dnspod.TTL},
			"format":      {"json"},
		},
		domain,
	)
	if err == nil && status.Status.Code == "1" {
		log.Println("新增域名解析", domain, "成功， IP 是", ipAddr)
		domain.UpdateStatus = config.UpdatedSuccess
	} else {
		log.Println("新增域名解析", domain, "失败， IP 是", ipAddr, "error message is", status.Status.Message)
		domain.UpdateStatus = config.UpdatedFail
	}
}

// update
func (dnspod *Dnspod) modify(result DnspodRecordListResp, domain *config.Domain, recordType, ipAddr string) {
	for _, record := range result.Records {
		// 相同的无需操作
		if record.Value == ipAddr {
			// log.Println("IP地址未发生变化，无需操作！", domain, "IP:", ipAddr)
			log.Printf("当前域名 %s 对应IP %s 未发生变化，无需操作。", domain, ipAddr)
			continue
		}

		status, err := dnspod.commonRequest(
			recordModifyURL,
			url.Values{
				"login_token": {dnspod.DNSConfig.ID + "," + dnspod.DNSConfig.Secret},
				"domain":      {domain.DomainName},
				"subDomain":   {domain.GetSubDomain()},
				"record_type": {recordType},
				"record_line": {"默认"},
				"record_id":   {record.ID},
				"value":       {ipAddr},
				"format":      {"json"},
				"ttl":         {dnspod.TTL},
			},
			domain,
		)
		if err == nil && status.Status.Code == "1" {
			log.Printf("更新域名 %s 成功,ip 是 %s", domain, ipAddr)
			domain.UpdateStatus = config.UpdatedSuccess
		} else {
			log.Printf("更新域名 %s 失败,ip 是 %s, error message is %s", domain, ipAddr, status.Status.Message)
			domain.UpdateStatus = config.UpdatedFail
		}
	}
}

// common
func (Dnspod *Dnspod) commonRequest(apiAddr string, values url.Values, domain *config.Domain) (status DnspodStatus, err error) {
	resp, err := http.PostForm(
		apiAddr,
		values,
	)
	err = util.GetHTTPResponse(resp, apiAddr, err, &status)

	return
}

// 获得域名记录
func (dnspod *Dnspod) getRecordList(domain *config.Domain, typ string) (result DnspodRecordListResp, err error) {
	values := url.Values{
		"login_token": {dnspod.DNSConfig.ID + "," + dnspod.DNSConfig.Secret},
		"domain":      {domain.DomainName},
		"record_type": {typ},
		"sub_domain":  {domain.GetSubDomain()},
		"format":      {"json"},
	}

	// if domain.SubDomain != "" {
	// 	values.Add("sub_domain", domain.SubDomain)
	// } else {
	// 	values.Add("Sub_domain", "@")
	// }

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.PostForm(
		recordListAPI,
		values,
		// url.Values{
		// 	"login_token": {dnspod.DNSConfig.ID + "," + dnspod.DNSConfig.Secret},
		// 	"domain":      {domain.DomainName},
		// 	"sub_domain":  {domain.SubDomain},
		// 	"record_type": {typ},
		// },
	)
	err = util.GetHTTPResponse(resp, recordListAPI, err, &result)

	return
}
