package web

import (
	"fmt"
	"io/ioutil"
	"myddns/config"
	"myddns/util"
	"net/http"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func Writing(writer http.ResponseWriter, request *http.Request) {
	data, err := Asset("static/pages/writing.html")
	if err != nil {
		// asset 没找到
	}
	tempFile := os.TempDir() + string(os.PathSeparator) + "writing.html"
	err = ioutil.WriteFile(tempFile, data, 0600)
	if err != nil {
		http.Error(writer, "Error writing temporary file", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(tempFile)
	if err != nil {
		fmt.Println("错误发生")
		fmt.Println(err)
		return
	}

	conf := &config.Config{}

	// 解析文件
	var configFile string = util.GetConfigFromFile()
	_, err = os.Stat(configFile)
	if err == nil {
		// 如果不是空
		byt, err := ioutil.ReadFile(configFile)
		if err == nil {
			err = yaml.Unmarshal(byt, conf)
			if err == nil {
				tmpl.Execute(writer, conf)
				return
			}
		}
	}

	// 默认值
	if conf.Ipv4.URL == "" {
		conf.Ipv4.URL = "http://api-ipv4.ip.sb/ip"
	}
	if conf.Ipv6.URL == "" {
		conf.Ipv6.URL = "http://api-ipv4.ip.sb/ip"
	}
	tmpl.Execute(writer, conf)
}
