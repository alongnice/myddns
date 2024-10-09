package dns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"strconv"
)

const (
	zonesAPI string = "https://api.cloudFlare.com/client/v4/zones"
)

// cloudFlare实现
type Cloudflare struct {
	DNSConfig config.DNSConfig
	Domains   config.Domains
	TTL       int
}

// CloudflareStatus 公共状态
type CloudflareStatus struct {
	Success  bool
	Messages []string
}

// 记录实体
type CloudflareRecord struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Proxied bool   `json:"proxied"`
	TTL     int    `json:"ttl"`
}

// 记录列表返回结果
type CloudflareRecordsResp struct {
	CloudflareStatus
	Result []CloudflareRecord
}

// cloudFlare zone 返回结果
type CloudflareZonesResp struct {
	CloudflareStatus
	Result []struct {
		ID     string
		Name   string
		Status string
		Paused string
	}
}

// 初始化
func (cf *Cloudflare) Init(conf *config.Config) {
	cf.DNSConfig = conf.DNS
	cf.Domains.GetNewIP(conf)
	if conf.TTL == "" {
		cf.TTL = 1
	} else {
		ttl, err := strconv.Atoi(conf.TTL)
		if err != nil {
			cf.TTL = 1
		} else {
			cf.TTL = ttl
		}
	}
}

// 添加或者更新IPv4/IPv6记录
func (cf *Cloudflare) AddUpdateDomainRecords() config.Domains {
	cf.AddUpdateIpvDomainRecords("A")
	cf.AddUpdateIpvDomainRecords("AAAA")
	return cf.Domains
}

func (cf *Cloudflare) AddUpdateIpvDomainRecords(recordType string) {
	ipAddr, domains := cf.Domains.GetNewIpResult(recordType)

	if ipAddr == "" {
		return
	}

	for _, domain := range domains {
		// get zone 获得域
		result, err := cf.getZones(domain)
		if err != nil || len(result.Result) != 1 {
			return
		}
		zoneID := result.Result[0].ID

		var records CloudflareRecordsResp
		// 最多取得50条记录
		cf.request(
			"GET",
			fmt.Sprintf(zonesAPI+"/%s/dns_records?type=%s&name=%s&per_page=50", zoneID, recordType, domain),
			nil,
			&records,
		)

		if records.Success && len(records.Result) > 0 {
			// 更新
			cf.modify(records, zoneID, domain, recordType, ipAddr)
			// 新增
			cf.create(zoneID, domain, recordType, ipAddr)
		}
	}
}
func (cf *Cloudflare) getZones(domain *config.Domain) (result CloudflareZonesResp, err error) {
	err = cf.request(
		"GET",
		fmt.Sprintf(zonesAPI+"?name=%s&status=%s&per_page=%s", domain.DomainName, "active", "50"),
		nil,
		&result,
	)
	return
}

// Update
func (cf *Cloudflare) modify(result CloudflareRecordsResp, zoneID string, domain *config.Domain, recordType string, ipAddr string) {
	for _, record := range result.Result {
		// 相同不修改
		if record.Content == ipAddr {
			log.Printf("当前域名 %s 对应IP %s 未发生变化，无需操作。", domain, ipAddr)
			continue
		}

		var status CloudflareStatus
		record.Content = ipAddr
		record.TTL = cf.TTL

		err := cf.request(
			"PUT",
			fmt.Sprintf(zonesAPI+"/%s/dns_records/%s", zoneID, record.ID),
			record,
			&status,
		)

		if err == nil && status.Success {
			log.Printf(("更新%s记录成功！ IP：%s"), recordType, ipAddr)
		} else {
			log.Printf(("更新%s记录失败！ IP：%s  Messages is %s"), recordType, ipAddr, status.Messages)
		}
	}
}

// create
func (cf *Cloudflare) create(zoneID string, domain *config.Domain, recordType string, ipAddr string) {
	record := &CloudflareRecord{
		Name:    domain.String(),
		Type:    recordType,
		Content: ipAddr,
		TTL:     cf.TTL,
		Proxied: false,
	}

	var status CloudflareStatus
	err := cf.request(
		"POST",
		fmt.Sprintf(zonesAPI+"/%s/dns_records", zoneID),
		record,
		&status,
	)

	if err == nil && status.Success {
		log.Printf(("新增%s记录成功！ IP：%s"), recordType, ipAddr)
		domain.UpdateStatus = config.UpdatedSuccess
	} else {
		log.Printf(("新增%s记录失败！ IP：%s  Messages is %s"), recordType, ipAddr, status.Messages)
		domain.UpdateStatus = config.UpdatedFail
	}
}

func (cf *Cloudflare) request(method string, url string, data interface{}, result interface{}) (err error) {
	jsonStr := make([]byte, 0)
	if data != nil {
		jsonStr, err = json.Marshal(data)
	}
	req, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer(jsonStr),
	)

	if err != nil {
		log.Println("http.NewRequest 失败： ", err)
		return
	}
	req.Header.Set("Authorization", "Bearer"+cf.DNSConfig.Secret)
	req.Header.Set("Content-Type", "application/json")

	// 创建一个http.Client，设置超时时间为30秒
	client := util.CreateHTTPClient()
	// 发送请求，获取响应
	resp, err := client.Do(req)
	err = util.GetHTTPResponse(resp, url, err, result)

	return nil
}
