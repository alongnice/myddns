package web

import (
	"myddns/config"
	"myddns/dns"
	"net/http"
	"strings"
)

// 保存
func Save(writer http.ResponseWriter, request *http.Request) {
	conf, _ := config.GetConfigCache()
	idNew := request.FormValue("DnsID")
	secretNew := request.FormValue("DnsSecret")

	idHide, secretHide := getHideIDSecret(&conf)

	if idNew != idHide {
		conf.DNS.ID = idNew
	}
	if secretNew != secretHide {
		conf.DNS.Secret = secretNew
	}

	conf.DNS.Name = request.FormValue("DnsName")
	// 从request中获取值，赋值给配置变量

	//从请求表格值设置IPv4配置值
	conf.Ipv4.Enable = request.FormValue("Ipv4Enable") == "on"
	conf.Ipv4.URL = strings.TrimSpace(request.FormValue("Ipv4Url"))
	conf.Ipv4.GetType = request.FormValue("Ipv4GetType")
	conf.Ipv4.NetInterface = request.FormValue("Ipv4NetInterface")
	conf.Ipv4.Domains = strings.Split(request.FormValue("Ipv4Domains"), "\r\n")

	conf.Ipv6.Enable = request.FormValue("Ipv6Enable") == "on"
	conf.Ipv6.GetType = request.FormValue("Ipv6GetType")
	conf.Ipv6.NetInterface = request.FormValue("Ipv6NetInterface")
	conf.Ipv6.URL = strings.TrimSpace(request.FormValue("Ipv6Url"))
	conf.Ipv6.Domains = strings.Split(request.FormValue("Ipv6Domains"), "\r\n")

	conf.Username = strings.TrimSpace(request.FormValue("Username"))
	conf.Password = request.FormValue("Password")

	conf.WebhookURL = strings.TrimSpace(request.FormValue("WebhookURL"))
	conf.WebhookRequestBody = strings.TrimSpace(request.FormValue("WebhookRequestBody"))

	// 保存到用户目录
	err := conf.SaveConfig()

	go dns.RunOnce()

	// 跳转
	if err == nil {
		writer.Write([]byte("保存成功"))
	} else {
		writer.Write([]byte(err.Error()))
	}
}
