package web

import (
	"embed"
	"fmt"
	"myddns/config"
	"net/http"
	"os"
	"strings"
	"text/template"
)

//go:embed writing.html
var writingEmbedFile embed.FS

const VersionEnv = "MYDDNS_VERSION"

type writtingData struct {
	config.Config
	Version string
}

func Writing(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFS(writingEmbedFile, "writing.html")
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

		tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
		return

	}

	// 默认值
	if conf.Ipv4.URL == "" {
		conf.Ipv4.URL = "https://myip.ipip.net, https://ddns.oray.com/checkip, https://ip.3322.net"
		conf.Ipv4.Enable = true
		conf.Ipv4.GetType = "url"
	}
	if conf.Ipv6.URL == "" {
		conf.Ipv6.URL = "https://api-ipv6.ip.sb/ip, https://speed.neu6.edu.cn/getIP.php, https://v6.ident.me"
		conf.Ipv6.GetType = "url"
	}
	if conf.DNS.Name == "" {
		conf.DNS.Name = "alidns"
	}
	// 默认禁止外部访问
	conf.NotAllowWanAccess = true

	tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
}

// 显示数目
const displayCount int = 3

func getHideIDSecret(conf *config.Config) (idHide string, secretHide string) {
	// webhook 显示所有ID
	if len(conf.DNS.ID) > displayCount && conf.DNS.Name != "callback" {
		idHide = conf.DNS.ID[:displayCount] + strings.Repeat("*", len(conf.DNS.ID)-displayCount)
	} else {
		idHide = conf.DNS.ID
	}
	// webhook 显示所有Secret
	if len(conf.DNS.Secret) > displayCount && conf.DNS.Name != "callback" {
		secretHide = conf.DNS.Secret[:displayCount] + strings.Repeat("*", len(conf.DNS.Secret)-displayCount)
	} else {
		secretHide = conf.DNS.Secret
	}
	return
}
