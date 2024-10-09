package dns

import (
	"bytes"
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"net/url"
)

const (
	alidnsEndpoint string = "https://alidns.aliyuncs.com"
)

// https://help.aliyun.com/document_detail/29776.html?spm=a2c4g.11186623.6.672.715a45caji9dMA
// 阿里云DNS实现
type Alidns struct {
	DNSConfig config.DNSConfig
	Domains   config.Domains
	TTL       string
}

type AlidnsSubDomainRecords struct {
	TotalCount    int
	DomainRecords struct {
		Record []struct {
			DomainName string
			RecordID   string
			Value      string
		}
	}
}
type AlidnsResp struct {
	RecordID  string
	RequestID string
}

// Init 初始化
func (ali *Alidns) Init(conf *config.Config) {
	ali.DNSConfig = conf.DNS
	ali.Domains.GetNewIP(conf)
	// 将原本解析域名的操作向下传递
	if conf.TTL == "" {
		ali.TTL = "600"
	} else {
		ali.TTL = conf.TTL
	}
}

// 添加或者更新IPv4/IPv6记录
func (ali *Alidns) AddUpdateDomainRecords() config.Domains {
	ali.AddUpdateIpvDomainRecords("A")
	ali.AddUpdateIpvDomainRecords("AAAA")
	return ali.Domains
}

func (ali *Alidns) AddUpdateIpvDomainRecords(recordType string) {
	ipAddr, domains := ali.Domains.GetNewIpResult(recordType)

	if ipAddr == "" {
		return
	}

	for _, domain := range domains {
		var record AlidnsSubDomainRecords
		// 获取当前域名信息
		params := url.Values{}
		params.Set("Action", "DescribeSubDomainRecords")
		params.Set("SubDomain", domain.GetFullDomain())
		params.Set("Type", recordType)
		err := ali.request(params, &record)

		if err != nil {
			return
		}

		if record.TotalCount > 0 {
			// 存在，更新
			ali.modify(record, domain, recordType, ipAddr)
		} else {
			// 不存在，创建
			ali.create(domain, recordType, ipAddr)
		}

	}
}

// 创建
func (ali *Alidns) create(domain *config.Domain, recordType string, ipAddr string) {
	params := url.Values{}
	params.Set("Action", "AddDomainRecord")
	params.Set("DomainName", domain.DomainName)
	params.Set("RR", domain.GetSubDomain())
	params.Set("Type", recordType)
	params.Set("Value", ipAddr)
	params.Set("TTL", ali.TTL)

	var result AlidnsResp
	err := ali.request(params, &result)

	if err == nil && result.RecordID != "" {
		log.Printf("新增域名解析 %s 成功！IP: %s", domain, ipAddr)
		domain.UpdateStatus = config.UpdatedSuccess
	} else {
		log.Printf("新增域名解析 %s 失败！", domain)
		domain.UpdateStatus = config.UpdatedFail
	}
}

// 修改
func (ali *Alidns) modify(record AlidnsSubDomainRecords, domain *config.Domain, recordType string, ipAddr string) {

	// 相同不修改
	if len(record.DomainRecords.Record) > 0 && record.DomainRecords.Record[0].Value == ipAddr {
		log.Printf("你的IP %s 没有变化, 域名 %s", ipAddr, domain)
		return
	}

	params := url.Values{}
	params.Set("Action", "UpdateDomainRecord")
	params.Set("RR", domain.GetSubDomain())
	params.Set("RecordId", record.DomainRecords.Record[0].RecordID)
	params.Set("Type", recordType)
	params.Set("Value", ipAddr)
	params.Set("TTL", ali.TTL)

	var result AlidnsResp
	err := ali.request(params, &result)

	if err == nil && result.RecordID != "" {
		log.Printf("更新域名解析 %s 成功！IP: %s", domain, ipAddr)
		domain.UpdateStatus = config.UpdatedSuccess
	} else {
		log.Printf("更新域名解析 %s 失败！", domain)
		domain.UpdateStatus = config.UpdatedFail
	}
}

// 定义request函数，参数为params和result，返回值为err
func (ali *Alidns) request(params url.Values, result interface{}) (err error) {

	// 使用util包中的AliyunSigner函数对params进行签名
	util.AliyunSigner(ali.DNSConfig.ID, ali.DNSConfig.Secret, &params)

	// 创建一个GET请求，请求地址为alidnsEndpoint，请求体为nil
	req, err := http.NewRequest(
		"GET",
		alidnsEndpoint,
		bytes.NewBuffer(nil),
	)
	// 将params编码为URL查询参数
	req.URL.RawQuery = params.Encode()

	// 如果创建请求失败，则打印错误信息并返回
	if err != nil {
		log.Println("http.NewRequest失败. Error: ", err)
		return
	}

	// 创建一个http.Client，设置超时时间为30秒
	client := util.CreateHTTPClient()
	// 发送请求，获取响应
	resp, err := client.Do(req)
	// 使用util包中的GetHTTPResponse函数处理响应，并将结果保存到result中
	err = util.GetHTTPResponse(resp, alidnsEndpoint, err, result)

	return
}
