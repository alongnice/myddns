package dns

import (
	"log"
	"myddns/config"

	alidnssdk "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	// "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	// 库后期将被弃用，暂时保留
)

// 阿里云DNS实现
type Alidns struct {
	client *alidnssdk.Client
	Domains
}

// Init 初始化
func (ali *Alidns) Init(conf *config.Config) {
	client, err := alidnssdk.NewClientWithAccessKey("cn-hangzhou", conf.DNS.ID, conf.DNS.Secret)
	if err != nil {
		log.Println("Ali dns 链接失败")
		// return false, false
	}
	ali.client = client

	// Ipv4
	ipv4Addr, err := conf.GetIpv4Addr()
	if err != nil {
		ali.Ipv4Addr = ipv4Addr
		ali.Ipv4Domains = ParseDomain(conf.Ipv4.Domains)
	}

	// Ipv6
	ipv6Addr, err := conf.GetIpv6Addr()
	if err != nil {
		ali.Ipv6Addr = ipv6Addr
		ali.Ipv6Domains = ParseDomain(conf.Ipv6.Domains)
	}

}

/* // AddRecord 添加记录
func (alidns *Alidns) AddRecord(conf *config.Config) (ipv4 bool, ipv6 bool) {

	ipv4Stat := addIpv4Record(client, conf)
	ipv6Stat := addIpv6Record(client, conf)
	return ipv4Stat, ipv6Stat
} */

func (ali *Alidns) AddUpdateIpv4DomainRecords() {
	ali.AddUpdateIpvDomainRecords("A")
}
func (ali *Alidns) AddUpdateIpv6DomainRecords() {
	ali.AddUpdateIpvDomainRecords("AAAA")
}

func (ali *Alidns) AddUpdateIpvDomainRecords(typ string) {
	typeName := "ipv4"
	ipAddr := ali.Ipv4Addr
	domains := ali.Ipv4Domains
	if typ == "AAAA" {
		typeName = "ipv6"
		ipAddr = ali.Ipv6Addr
		domains = ali.Ipv6Domains
	}
	if ipAddr == "" {
		return
	}

	existReq := alidnssdk.CreateDeleteSubDomainRecordsRequest()
	existReq.Type = typ

	for _, dom := range domains {
		existReq.SubDomain = dom.SubDomain + "." + dom.DomainName
		rep, err := ali.client.DescribeSubDomainRecords(existReq)
		if err != nil {
			log.Println(err)
		}
		if rep.TotalCount > 0 {
			// 更新
			if rep.DomainRecords.Record[0].Value != ipAddr {
				request := alidnssdk.CreateUpdateDomainRecordRequest()
				request.Scheme = "https"
				request.Value = ipAddr
				request.Type = typ
				request.RR = dom.SubDomain
				request.RecordId = rep.DomainRecords.Record[0].RecordId

				_, err := ali.client.UpdateDomainRecord(request)
				if err != nil {
					log.Println("更新ipv4记录错误！ Domain: ", dom, " ip: ", ipAddr, "ERROR")
				} else {
					log.Println("更新ipv4记录成功！ Domain: ", dom, " ip: ", ipAddr)
				}
				if rep.TotalCount > 1 {
					log.Println(dom, "存在多条记录，只会更新第一条")
				}
			} else {
				// 新增
				request := alidnssdk.CreateAddDomainRecordRequest()
				request.Scheme = "https"
				request.Value = ipAddr
				request.Type = typ
				request.RR = dom.SubDomain
				request.DomainName = dom.DomainName

				_, err := ali.client.AddDomainRecord(request)
				if err != nil {
					log.Println("添加", typeName, " 错误，Domain: ", dom, " ip: ", ipAddr)
				} else {
					log.Println("添加", typeName, " 成功，Domain: ", dom, " ip: ", ipAddr)

				}
			}
		}
	}
}
