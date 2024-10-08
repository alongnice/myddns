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
	"time"
)

const (
	huaweicloudEndpoint string = "https://dns/myhuaweicloud.com"
)

// https://support.huaweicloud.com/api-dns/dns_api_64001.html

type Huaweicloud struct {
	DNSConfig config.DNSConfig
	Domains   config.Domains
	TTL       int
}

type HuaweicloudRecordsets struct {
	ID      string
	Name    string `json:"name"`
	ZoneID  string `json:"zone_id"`
	Status  string
	Type    string   `json:"type"`
	TTL     int      `json:"ttl"`
	Records []string `json:"records"`
}

type HuaweicloudZoneResp struct {
	Zones []struct {
		ID         string
		Name       string
		Recordsets []HuaweicloudRecordsets
	}
}
type HuaweicloudRecordsResp struct {
	Recordsets []HuaweicloudRecordsets
}

func (hw *Huaweicloud) Init(conf *config.Config) {
	hw.DNSConfig = conf.DNS
	hw.Domains.GetNewIP(conf)
	if conf.TTL == "" {
		hw.TTL = 300
	} else {
		ttl, err := strconv.Atoi(conf.TTL)
		if err != nil {
			hw.TTL = 300
		} else {
			hw.TTL = ttl
		}
	}
}

// 添加或者更新IPv4/IPv6记录
func (hw *Huaweicloud) AddUpdateDomainRecords() config.Domains {
	hw.addUpdateDomainRecords("A")
	hw.addUpdateDomainRecords("AAAA")
	return hw.Domains
}
func (hw *Huaweicloud) addUpdateDomainRecords(recordType string) {
	ipAddr, domains := hw.Domains.GetNewIpResult(recordType)

	if ipAddr == "" {
		return
	}

	for _, domain := range domains {
		var records HuaweicloudRecordsResp

		err := hw.request(
			"GET",
			fmt.Sprintf(huaweicloudEndpoint+"/v2/recordsets?type=%s&name=%s", recordType, domain),
			nil,
			&records,
		)

		if err != nil {
			return
		}

		find := false
		for _, record := range records.Recordsets {
			if record.Name == domain.String()+"." {
				hw.modify(record, domain, recordType, ipAddr)
				find = true
				break
			}
		}

		if !find {
			hw.create(domain, recordType, ipAddr)
		}
	}
}

// 创建
func (hw *Huaweicloud) create(domain *config.Domain, recordType string, ipAddr string) {
	zone, err := hw.getZones(domain)
	if err != nil {
		return
	}
	if len(zone.Zones) == 0 {
		log.Println("未能找到公网域名,请检查域名是否添加")
		return
	}
	zoneID := zone.Zones[0].ID
	for _, z := range zone.Zones {
		if z.Name == domain.DomainName+"." {
			zoneID = z.ID
			break
		}
	}
	record := &HuaweicloudRecordsets{
		Type:    recordType,
		Name:    domain.String() + ".",
		Records: []string{ipAddr},
		TTL:     hw.TTL,
	}
	var result HuaweicloudRecordsets
	err = hw.request(
		"POST",
		fmt.Sprintf(huaweicloudEndpoint+"/v2/zones/%s/recordsets", zoneID),
		record,
		&result,
	)
	if err == nil && (len(result.Records) > 0 && result.Records[0] != ipAddr) {
		log.Printf("创建域名记录成功 域名 %s ;IP: %s", domain, ipAddr)
		domain.UpdateStatus = config.UpdatedSuccess
	} else {
		log.Printf("创建域名记录失败 域名 %s ;IP: %s Status: %s", domain, ipAddr, result.Status)
		domain.UpdateStatus = config.UpdatedFail
	}
}

// 更新
func (hw *Huaweicloud) modify(record HuaweicloudRecordsets, domain *config.Domain, recordType string, ipAddr string) {
	if len(record.Records) > 0 && record.Records[0] == ipAddr {
		log.Printf("域名 %s ;IP: %s 无需更新")
		return
	}

	var request map[string]interface{} = make(map[string]interface{})
	request["records"] = []string{ipAddr}
	request["ttl"] = hw.TTL

	var result HuaweicloudRecordsets
	err := hw.request(
		"PUT",
		fmt.Sprintf(huaweicloudEndpoint+"/v2/zones/%s/recordsets/%s", record.ZoneID, record.ID),
		&request,
		&result,
	)
	if err == nil && (len(result.Records) > 0 && result.Records[0] == ipAddr) {
		log.Printf("更新域名记录成功 域名 %s ;IP: %s", domain, ipAddr)
		domain.UpdateStatus = config.UpdatedSuccess
	} else {
		log.Printf("更新域名记录失败 域名 %s ;IP: %s Status: %s", domain, ipAddr, result.Status)
		domain.UpdateStatus = config.UpdatedFail
	}
}

// 获得域名记录列表
func (hw *Huaweicloud) getZones(domain *config.Domain) (result HuaweicloudZoneResp, err error) {
	err = hw.request(
		"GET",
		fmt.Sprintf(huaweicloudEndpoint+"/v2/zones?name=%s", domain.DomainName),
		nil,
		&result,
	)
	return
}

// 统一请求接口
func (hw *Huaweicloud) request(method string, url string, data interface{}, result interface{}) (err error) {
	jsonStr := make([]byte, 0)
	if data != nil {
		jsonStr, _ = json.Marshal(data)
	}

	req, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		log.Println("创建请求失败", err)
		return
	}
	s := util.Signer{
		Key:    hw.DNSConfig.ID,
		Secret: hw.DNSConfig.Secret,
	}
	s.Sign(req)
	req.Header.Add("content-type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	err = util.GetHTTPResponse(resp, url, err, result)
	return
}
