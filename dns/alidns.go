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

// func (ali *Alidns) init(conf *config.Config) {
// 	ali.Init(conf)
// }

// Init 初始化
func (ali *Alidns) Init(conf *config.Config) {
	client, err := alidnssdk.NewClientWithAccessKey("cn-hangzhou", conf.DNS.ID, conf.DNS.Secret)
	if err != nil {
		log.Println("Ali dns 链接失败")
		// return false, false
	}
	ali.client = client

	ali.Domains.ParseDomain(conf)
	// 将原本解析域名的操作向下传递
}

// 添加或者更新IPv4记录
func (ali *Alidns) AddUpdateIpv4DomainRecords() {
	ali.AddUpdateIpvDomainRecords("A")
}

// 添加或者更新IPv6记录
func (ali *Alidns) AddUpdateIpv6DomainRecords() {
	ali.AddUpdateIpvDomainRecords("AAAA")
}

func (ali *Alidns) AddUpdateIpvDomainRecords(recordType string) {
	typeName := "ipv4"
	ipAddr := ali.Ipv4Addr
	domains := ali.Ipv4Domains
	if recordType == "AAAA" {
		typeName = "ipv6"
		ipAddr = ali.Ipv6Addr
		domains = ali.Ipv6Domains
	}
	if ipAddr == "" {
		return
	}

	existReq := alidnssdk.CreateDescribeSubDomainRecordsRequest()
	existReq.Type = recordType

	for _, domain := range domains {
		existReq.SubDomain = domain.SubDomain + "." + domain.DomainName
		rep, err := ali.client.DescribeSubDomainRecords(existReq)
		if err != nil {
			log.Println(err)
		}
		if rep.TotalCount > 0 {
			// 更新
			for _, record := range rep.DomainRecords.Record {
				if record.Value == ipAddr {
					log.Printf("当前域名 %s 对应IP %s 未发生变化，无需操作。", domain, ipAddr)
					continue
				}
				request := alidnssdk.CreateUpdateDomainRecordRequest()
				request.Scheme = "https"
				request.Value = ipAddr
				request.Type = recordType
				request.RR = domain.SubDomain
				request.RecordId = record.RecordId

				updateResp, err := ali.client.UpdateDomainRecord(request)
				if err != nil || !updateResp.BaseResponse.IsSuccess() {
					log.Println("更新ipv4记录错误！", typeName, " Domain: ", domain, " ip: ", ipAddr, "ERROR", err, "Response is", updateResp.GetHttpContentString())
				} else {
					log.Println("更新ipv4记录成功！", typeName, " Domain: ", domain, " ip: ", ipAddr)
				}
				// if rep.TotalCount > 1 {
				// 	log.Println(typeName, dom, "存在多条记录，只会更新第一条")
				// }
			}
		} else {
			// 新增
			request := alidnssdk.CreateAddDomainRecordRequest()
			request.Scheme = "https"
			request.Value = ipAddr
			request.Type = recordType
			request.RR = domain.SubDomain
			request.DomainName = domain.DomainName

			createResp, err := ali.client.AddDomainRecord(request)
			if err != nil {
				log.Println("添加", typeName, " 错误，Domain: ", domain, " ip: ", ipAddr, "error is ", err, "createResp is ", createResp.GetHttpContentString())
			} else {
				log.Println("添加", typeName, " 成功，Domain: ", domain, " ip: ", ipAddr)

			}
		}
	}
}
