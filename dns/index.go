package dns

import (
	"log"
	"myddns/config"
	"strings"
	"time"
)

// DNS 接口 添加初始化和更新的方法
type DNS interface {
	// AddRecord(conf *config.Config) (ipv4 bool, ipv6 bool)
	Init(conf *config.Config)
	AddUpdateIpv4DomainRecords()
	// 添加或更新IPv4记录
	AddUpdateIpv6DomainRecords()
	// 添加或更新IPv6记录
}

// ipv4,ipv6的域
type Domains struct {
	Ipv4Addr    string
	Ipv4Domains []*Domain
	Ipv6Addr    string
	Ipv6Domains []*Domain
}

// Domain 域名实体对象包含 域名和子域名 以及是否可用
type Domain struct {
	DomainName string
	SubDomain  string
	Exist      bool
}

// 创建域名实体对象的构造（字符串处理）
func (d Domain) String() string {
	if d.SubDomain != "" {
		return d.SubDomain + "." + d.DomainName
	}
	return d.DomainName

}

// 获取全部子域名
func (d Domain) GetFullDomain() string {
	if d.SubDomain != "" {
		return d.SubDomain + "." + d.DomainName
	}
	return "@" + "." + d.DomainName
}

// GetSubDomain 获得子域名，为空返回@
// 阿里云，dnspod需要
func (d Domain) GetSubDomain() string {
	if d.SubDomain != "" {
		return d.SubDomain
	}
	return "@"
}

// runOnce
func RunOnce() {
	conf := &config.Config{}
	err := conf.InitConfigFromFile()
	if err != nil {
		return
	}

	var dnsSelected DNS
	switch conf.DNS.Name {
	case "alidns":
		dnsSelected = &Alidns{}
	case "dnspod":
		dnsSelected = &Dnspod{}
	case "cloudflare":
		dnsSelected = &Cloudflare{}
	default:
		dnsSelected = &Alidns{}

	}
	dnsSelected.Init(conf)
	dnsSelected.AddUpdateIpv4DomainRecords()
	dnsSelected.AddUpdateIpv6DomainRecords()
}

// RunTimer 定时运行
func RunTimer() {
	for {
		RunOnce()
		time.Sleep(time.Second * time.Duration(5))
	}
}

// ParseDomain 解析域名,传入域名数组 返回域名实体数组
func (domains *Domains) ParseDomain(conf *config.Config) {
	// ipv4
	ipv4Addr := conf.GetIpv4Addr()
	if ipv4Addr != "" {
		domains.Ipv4Addr = ipv4Addr
		domains.Ipv4Domains = ParseDomainInnerchan(conf.Ipv4.Domains)
	}

	// ipv6
	ipv6Addr := conf.GetIpv6Addr()
	if ipv6Addr != "" {
		domains.Ipv6Addr = ipv6Addr
		domains.Ipv6Domains = ParseDomainInnerchan(conf.Ipv6.Domains)
	}
}

func ParseDomainInnerchan(domainArr []string) (domains []*Domain) {
	// 解析域名
	for _, domainStr := range domainArr {
		domainStr = strings.Trim(domainStr, " ")
		// 去除空格
		if domainStr != "" {
			// 遍历域名数组
			domain := &Domain{}
			// 构建一个域名实体对象
			sp := strings.Split(domainStr, ".")
			// 切片分割，进行异常判断
			length := len(sp)
			if length <= 1 { // 错误情况
				log.Println(domainStr, "域名不正确")
				continue
			} else if length == 2 { //单个域名情况
				domain.DomainName = domainStr
			} else { // >=3
				domain.DomainName = sp[length-2] + "." + sp[length-1]
				// 子域名
				domain.SubDomain = domainStr[:len(domainStr)-len(domain.DomainName)-1]
				// 判断是否可用
			}
			domains = append(domains, domain)
			// 添加到域名实体数组
		}
	}
	return
}
