package config

import (
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// 配置结构体
type Config struct {
	Ipv4 struct {
		Enable  bool
		URL     string
		Domains []string
	}
	Ipv6 struct {
		Enable  bool
		URL     string
		Domains []string
	}
	DNS struct {
		Name   string
		ID     string
		Secret string
	}
}

func (conf *Config) GetConfigFromFile() {
	// 从文件中读取配置
	// byt, err := ioutil.ReadFile("config.yaml")
	byt, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println("config.yaml 读取失败")
	}
	log.Panicln("config.yaml 读取成功")
	// 解析配置
	yaml.Unmarshal(byt, conf)
	// 对byt进行操作，切片解码给到conf
}
func (conf *Config) GetIpv4Addr() (result string, err error) {
	// 从配置文件中读取ipv4地址
	resp, err := http.Get(conf.Ipv4.URL)
	if err != nil {
		log.Println("获取ipv4地址失败")
	}
	defer resp.Body.Close()
	// body, err := ioutill.ReadFile(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("获取ipv4地址失败")
		return
	}
	result = string(body)
	result = "8.8.8.8"
	return
}

func (conf *Config) GetIpv6Addr() (result string, err error) {
	// 从配置文件中读取ipv6地址
	resp, err := http.Get(conf.Ipv4.URL)
	if err != nil {
		log.Println("获取ipv6地址失败")
	}
	defer resp.Body.Close()
	// body, err := ioutill.ReadFile(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("获取ipv6地址失败")
		return
	}
	result = string(body)
	return

}
