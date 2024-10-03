package web

import (
	"fmt"
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"strings"
	"text/template"
)

func Writing(writer http.ResponseWriter, request *http.Request) {
	tempPath, err := util.GetStaticResourcePath("static/pages/writing.html")
	if err != nil {
		log.Println(tempPath, "Asset 没找到.")
		return
	}
	tmpl, err := template.ParseFiles(tempPath)
	if err != nil {
		fmt.Println("Error 发生..")
		fmt.Println(err)
		return
	}

	conf, err := config.GetConfigCache()

	if err == nil {
		// 隐藏真实的ID,Secret
		idHide, secretHide := getHideIDSecret(&conf)
		conf.DNS.ID = idHide
		conf.DNS.Secret = secretHide

		tmpl.Execute(writer, &conf)
		return

	}

	// 默认值
	if conf.Ipv4.URL == "" {
		conf.Ipv4.URL = "https://myip.ipip.net"
		conf.Ipv4.Enable = true
	}
	if conf.Ipv6.URL == "" {
		conf.Ipv6.URL = "https://api-ipv6.ip.sb/ip"
	}
	if conf.DNS.Name == "" {
		conf.DNS.Name = "alidns"
	}

	tmpl.Execute(writer, conf)
}

// 显示数目
const displayCount int = 3

func getHideIDSecret(conf *config.Config) (idHide string, secretHide string) {
	// webhook 显示所有ID
	if len(conf.DNS.ID) > displayCount && conf.DNS.Name == "webhook" {
		idHide = conf.DNS.ID[:displayCount] + strings.Repeat("*", len(conf.DNS.ID)-displayCount)
	}else{
		idHide = conf.DNS.ID
	}
	// webhook 显示所有Secret
	if len(conf.DNS.Secret) > displayCount && conf.DNS.Name == "webhook"{
		secretHide = conf.DNS.Secret[:displayCount] + strings.Repeat("*", len(conf.DNS.Secret)-displayCount)
	}else{
		secretHide = conf.DNS.Secret
	}
	return
}
