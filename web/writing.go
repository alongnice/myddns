package web

import (
	"fmt"
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"text/template"
	"strings"
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

	conf := &config.Config{}
	err = conf.InitConfigFromFile()

	if err == nil {
		// 隐藏真实的ID,Secret
		idHide, secretHide := getHideSecret(conf)
		conf.DNS.ID = idHide
		conf.DNS.Secret = secretHide

		tmpl.Execute(writer, conf)
		return 

	}

	// 默认值
	if conf.Ipv4.URL == "" {
		conf.Ipv4.URL = "http://api-ipv4.ip.sb/ip"
		conf.Ipv4.Enable = true
	}
	if conf.Ipv6.URL == "" {
		conf.Ipv6.URL = "http://api-ipv4.ip.sb/ip"
	}
	if conf.DNS.Name == "" {
		conf.DNS.Name = "alidns"
	}

	tmpl.Execute(writer, conf)
}


// 显示数目
const displayCount = 3
func getHideSecret(conf *config.Config) (idHide string, secretHide string) {
	if len(conf.DNS.Secret) > displayCount{
		idHide = conf.DNS.ID[:displayCount] + strings.Repeat("*", len(conf.DNS.ID) - displayCount)
	}
	if len(conf.DNS.Secret) > displayCount{
	    secretHide = conf.DNS.Secret[:displayCount] + strings.Repeat("*", len(conf.DNS.Secret) - displayCount)
	}
	return 
}