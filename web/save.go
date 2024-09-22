package web

import (
	"io/ioutil"
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

// 保存
func Save(writer http.ResponseWriter, request *http.Request) {
	conf := &config.Config{}
	// 创建一个配置变量

	conf.DNS.Name = request.FormValue("DnsName")
	// 从request中获取值，赋值给配置变量
	conf.DNS.ID = request.FormValue("DnsID")
	conf.DNS.Secret = request.FormValue("DnsSecret")

	conf.Ipv4.Enable = request.FormValue("Ipv4Enable") == "on"
	conf.Ipv4.URL = request.FormValue("Ipv4Url")
	conf.Ipv4.Domains = strings.Split(request.FormValue("Ipv4Domains"), "\r\n")
	// 对多个结果进行切片,然后复制

	conf.Ipv6.Enable = request.FormValue("Ipv6Enable") == "on"
	conf.Ipv6.URL = request.FormValue("Ipv6Url")
	conf.Ipv4.Domains = strings.Split(request.FormValue("Ipv6Domains"), "\r\n")

	// 保存到用户目录
	util.GetConfigFromFile()
	// 打开配置文件
	byt, err := yaml.Marshal(conf)
	// 将配置变量转换为yaml格式
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile(util.GetConfigFromFile(), byt, 0644)
	// 写入配置文件 从byt中读取 文件权限644

	// 跳转
	http.Redirect(writer, request, "/?saveOk=true", http.StatusFound)
	// 重定向到首页
}
