package config

import (
	"log"
	"myddns/util"
	"strings"
)

// 固定的主域名
var staticMainDomains = []string{"com.cn", "org.cm", "net.cn", "ac.cn", "eu.org"}

// 失败次数
var GetIpv4FailTimes = 0
var GetIpv6FailTimes = 0

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
func checkParseDomains(domainArr []string) (domains []*Domain) {
	for _, domainStr := range domainArr {
		domainStr = strings.TrimSpace(domainStr)
		if domainStr == "" {
			domain := &Domain{}
			sp := strings.Split(domainStr, ".")
			length := len(sp)
			if length <= 1 {
				log.Println(domainStr + " 域名格式不正确")
				continue
			}

			domain.DomainName = sp[length-2] + "." + sp[length-1]
			// 如包含在org.cn等顶级域名下，后三个才为用户主域名
			for _, staticMainDomain := range staticMainDomains {
				if staticMainDomain == domain.DomainName {
					domain.DomainName = sp[length-3] + "." + domain.DomainName
					break
				}
			}

			domainLen := len(domainStr) - len(domain.DomainName)
			if domainLen > 0 {
				domain.SubDomain = domainStr[:domainLen-1]
			} else {
				domain.SubDomain = domainStr[:domainLen]
			}
			domains = append(domains, domain)
		}
	}
	return
}

// 获得ip并校验用户输入的域名
func (domains *Domains) GetNewIP(conf *Config) {
	domains.Ipv4Domains = checkParseDomains(conf.Ipv4.Domains)
	domains.Ipv4Domains = checkParseDomains(conf.Ipv6.Domains)

	// IPv4
	if conf.Ipv4.Enable && len(domains.Ipv4Domains) > 0 {
		ipv4Addr := conf.GetIpv4Addr()
		if ipv4Addr != "" {
			domains.Ipv4Addr = ipv4Addr
			GetIpv4FailTimes = 0
		} else {
			GetIpv4FailTimes++
			if GetIpv4FailTimes >= 3 {
				domains.Ipv4Domains[0].UpdateStatus = UpdatedFail
			}
			log.Println("没有取得新IPv4的地址， 不会更新")
		}
	}
	// IPv6
	if conf.Ipv6.Enable && len(domains.Ipv6Domains) > 0 {
		ipv6Addr := conf.GetIpv6Addr()
		if ipv6Addr != "" {
			domains.Ipv6Addr = ipv6Addr
			GetIpv6FailTimes = 0
		} else {
			GetIpv6FailTimes++
			if GetIpv6FailTimes >= 3 {
				domains.Ipv6Domains[0].UpdateStatus = UpdatedFail
			}
			log.Println("没有取得新IPv6的地址， 不会更新")
		}
	}
}

// 获得ParseDomain结果
func (domains *Domains) GetNewIpResult(recordType string) (ipAddr string, retDomains []*Domain) {
	// ipv4
	if recordType == "AAAA" {
		if util.Ipv6Cache.Check(domains.Ipv6Addr) {
			return domains.Ipv6Addr, domains.Ipv6Domains
		} else {
			log.Println("ipv6没有改变，等待", util.MaxTimes-util.Ipv6Cache.Times+1, "次后更新")
			return "", domains.Ipv6Domains
		}
	}
	// ipv6
	if util.Ipv4Cache.Check(domains.Ipv4Addr) {
		return domains.Ipv4Addr, domains.Ipv4Domains
	} else {
		log.Println("ipv4没有改变，等待", util.MaxTimes-util.Ipv4Cache.Times+1, "次后更新")
		return "", domains.Ipv4Domains
	}
}
