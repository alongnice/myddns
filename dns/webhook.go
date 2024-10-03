package dns

import (
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"strings"
	"time"
)

type Webhook struct {
	DNSConfig config.DNSConfig
	Domains
}

func (wh *Webhook) Init(conf *config.Config) {
	wh.DNSConfig = conf.DNS
	wh.Domains.ParseDomain(conf)
}

var beforeDomains Domains

func (wh *Webhook) AddUpdateDomainRecords() {
	if beforeDomains.Ipv4Addr == "" && beforeDomains.Ipv6Addr == "" {
		log.Println("暂时没有以前的IP地址，无需更新DNS记录,等待下一轮检测")
	} else {
		if beforeDomains.Ipv4Addr != wh.Domains.Ipv4Addr || beforeDomains.Ipv6Addr != wh.Domains.Ipv6Addr {
			method := "GET"
			postPara := ""
			if wh.DNSConfig.Secret != "" {
				method = "POST"
				postPara = wh.replacePara(wh.DNSConfig.Secret)
			}

			req, err := http.NewRequest(method, wh.replacePara(wh.DNSConfig.ID), strings.NewReader(postPara))

			clt := http.Client{}
			clt.Timeout = 30 * time.Second
			resp, err := clt.Do(req)
			body, err := util.GetHTTPResponseOrg(resp, wh.replacePara(wh.DNSConfig.ID), err)
			if err == nil {
				if wh.Domains.Ipv4Addr != "" {
					log.Println("更新IPv4地址为:", wh.Domains.Ipv4Addr)
				}
				if wh.Domains.Ipv6Addr != "" {
					log.Println("更新IPv4地址为:", wh.Domains.Ipv4Addr)
				}
				log.Printf("返回数据: %s", string(body))
			} else {
				log.Println("webhook调用失败,下次重新调用")
				return
			}
		} else {
			if wh.Domains.Ipv4Addr != "" {
				log.Println("你的IP: %s 没有变化，无需更新DNS记录", wh.Domains.Ipv4Addr)
			}
			if wh.Domains.Ipv6Addr != "" {
				log.Println("你的IP: %s 没有变化，无需更新DNS记录", wh.Domains.Ipv6Addr)
			}
		}
	}
	beforeDomains = wh.Domains
}

// replacePara函数用于替换字符串中的参数
func (wh *Webhook) replacePara(orgPara string) (newPara string) {
	// 将字符串中的{ipv4New}替换为wh.Domains.Ipv4Addr
	orgPara = strings.ReplaceAll(orgPara, "{ipv4New}", wh.Domains.Ipv4Addr)
	// 将字符串中的{ipv4Old}替换为beforeDomains.Ipv4Addr
	orgPara = strings.ReplaceAll(orgPara, "{ipv4Old}", beforeDomains.Ipv4Addr)
	// 将字符串中的{ipv6New}替换为wh.Domains.Ipv6Addr
	orgPara = strings.ReplaceAll(orgPara, "{ipv6New}", wh.Domains.Ipv6Addr)
	// 将字符串中的{ipv6Old}替换为beforeDomains.Ipv6Addr
	orgPara = strings.ReplaceAll(orgPara, "{ipv6Old}", beforeDomains.Ipv6Addr)
	// 将字符串中的{ipv4Domains}替换为wh.Domains.Ipv4Domains的字符串表示
	orgPara = strings.ReplaceAll(orgPara, "{ipv4Domains}", getDomainStr(wh.Domains.Ipv4Domains))
	// 将字符串中的{ipv6Domains}替换为wh.Domains.Ipv6Domains的字符串表示
	orgPara = strings.ReplaceAll(orgPara, "{ipv6Domains}", getDomainStr(wh.Domains.Ipv6Domains))

	return orgPara
}

// 根据传入的Domain指针数组，返回一个字符串
func getDomainStr(domains []*Domain) string {
	// 初始化字符串
	str := ""
	// 遍历Domain指针数组
	for i, v46 := range domains {
		// 将每个Domain指针转换为字符串并添加到str中
		str += v46.String()
		// 如果不是最后一个元素，则在字符串末尾添加逗号
		if i < len(domains)-1 {
			str += ","
		}
	}
	// 返回字符串
	return str
}
