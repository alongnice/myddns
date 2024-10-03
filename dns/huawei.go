package dns

import(
	"myddns/config"
	"myddns/util"
	"fmt"
	"log"
	"net/http"
	"time"
	"bytes"
	"encoding/json"
)

const (
	huaweicloudEndpoint string = "https://dns/myhuaweicloud.com"
)

type Huaweicloud struct{
	DNSConfig config.DNSConfig
	Domains
}

type HuaweicloudRecordsets struct{
    ID string
	Name string `json:"name"`
	ZoneID string `json:"zone_id"`
	Status string
	Type string `json:"type"`
	Records []string `json:"records"`
}

type HuaweicloudZoneResp struct{
	Zones []struct{
		ID string
		Name string
		Recordsets []HuaweicloudRecordsets
	}
}
type HuaweicloudRecordsResp struct{
	Recordsets []HuaweicloudRecordsets
}

func (hw *Huaweicloud) Init(conf *config.Config){
	hw.DNSConfig = conf.DNS
	hw.Domains.ParseDomain(conf)
}
func (hw *Huaweicloud) AddUpdateIpv4DomainRecords(){
	hw.AddUpdateDomainRecords("A")
}
func (hw *Huaweicloud) AddUpdateIpv6DomainRecords(){
	hw.AddUpdateDomainRecords("AAAA")
}
func (hw *Huaweicloud) AddUpdateDomainRecords(recordType string){
	ipAddr := hw.Ipv4Addr
	domains := hw.Ipv4Domains
	if recordType == "AAAA"{
	    ipAddr = hw.Ipv6Addr
		domains = hw.Domains.Ipv6Domains
	}
	if ipAddr == ""{
	    return
	}

	for _, domain := range domains{
	    var records HuaweicloudRecordsResp

		err := hw.request(
			"GET",
			fmt.Sprintf(huaweicloudEndpoint+"/v2/recordsets?type=%s&name=%s", recordType, domain),
			nil,
			&records,
		)

		if err != nil{
		    return 
		}

		find := false
		for _, record := range records.Recordsets{
			if record.Name == domain.String()+"."{
			    hw.modify(record, domain, recordType, ipAddr)
				find = true
				break
			}
		}

		if !find{
		    hw.create(domain, recordType, ipAddr)
		}
	}
}


// 创建
func (hw *Huaweicloud) create(domain *Domain, recordType string, ipAddr string){
	zone, err := hw.getZones(domain)
	if err != nil || len(zone.Zones) == 0{
		log.Println("未能找到公网域名,请检查域名是否添加")
	    return
	}
	zoneID := zone.Zones[0].ID
	for _,z := range zone.Zones{
	    if z.Name == domain.DomainName+"."{
			zoneID = z.ID
			break
		}
	}
	record := &HuaweicloudRecordsets{
		Type: recordType,
		Name: domain.String()+".",
		Records: []string{ipAddr},
	}
	var result HuaweicloudRecordsets
	err = hw.request(
		"POST",
		fmt.Sprintf(huaweicloudEndpoint+"/v2/zones/%s/recordsets", zoneID),
		record,
		&result,
	)
	if err == nil && (len(result.Records) > 0 && result.Records[0] != ipAddr){
		log.Printf("创建域名记录成功 域名 %s ;IP: %s", domain, ipAddr)
	}else{
		log.Printf("创建域名记录失败 域名 %s ;IP: %s Status: %s", domain, ipAddr, result.Status)
	}
}
// 更新
func (hw *Huaweicloud) modify(record HuaweicloudRecordsets, domain *Domain, recordType string, ipAddr string){
	if len(record.Records) > 0 && record.Records[0] == ipAddr{
		log.Printf("域名 %s ;IP: %s 无需更新")
		return
	}

	var request map[string]interface{} = make(map[string]interface{})
	request["records"] = []string{ipAddr}

	var result HuaweicloudRecordsets
	err := hw.request(
		"PUT",
		fmt.Sprintf(huaweicloudEndpoint+"/v2/zones/%s/recordsets/%s", record.ZoneID, record.ID),
		&request,
		&result,
	)
	if err == nil && (len(result.Records) > 0 && result.Records[0] == ipAddr){
		log.Printf("更新域名记录成功 域名 %s ;IP: %s", domain, ipAddr)
	}else{
		log.Printf("更新域名记录失败 域名 %s ;IP: %s Status: %s", domain, ipAddr, result.Status)
	}
}
// 获得域名记录列表
func (hw *Huaweicloud) getZones(domain *Domain)(result HuaweicloudZoneResp, err error){
	err = hw.request(
		"GET",
		fmt.Sprintf(huaweicloudEndpoint+"/v2/zones?name=%s", domain.DomainName),
		nil,
		&result,
	)
	return
}
// 统一请求接口
func (hw *Huaweicloud) request(method string, url string, data interface{}, result interface{}) (err error){
	jsonStr := make([]byte, 0)
	if data != nil{
		jsonStr, _ = json.Marshal(data)
	}

	req, err := http.NewRequest(
		method, 
		url, 
		bytes.NewBuffer(jsonStr),
	)
	if err != nil{
	    log.Println("创建请求失败", err)
		return
	}
	s := util.Signer{
		Key: hw.DNSConfig.ID,
		Secret: hw.DNSConfig.Secret,
	}
	s.Sign(req)
	req.Header.Add("content-type", "application/json")

	clt := http.Client{}
	clt.Timeout = 1 * time.Minute
	resp, err := clt.Do(req)
	err = util.GetHTTPResponse(resp, url, err, result)
	return
}