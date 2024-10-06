package dns

import (
	"myddns/config"
	"time"
)

// DNS 接口 添加初始化和更新的方法
type DNS interface {
	// AddRecord(conf *config.Config) (ipv4 bool, ipv6 bool)
	Init(conf *config.Config)
	// 添加或更新 IPV4/IPV6 记录
	AddUpdateDomainRecords() (domains config.Domains)
}

// runOnce
func RunOnce() {
	conf, err := config.GetConfigCache()
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
	case "huaweicloud":
		dnsSelected = &Huaweicloud{}
	default:
		dnsSelected = &Alidns{}
	}
	dnsSelected.Init(&conf)

	domains := dnsSelected.AddUpdateDomainRecords()
	domains.ExecWebhook(&conf)
}

// RunTimer 定时运行
func RunTimer() {
	for {
		RunOnce()
		time.Sleep(time.Second * time.Duration(5))
	}
}
