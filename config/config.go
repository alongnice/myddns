package config

import (
	"io"
	"log"
	"myddns/util"
	"net/http"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

// ipv4reg ipv4的正则
const Ipv4Reg = `((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])`

// ipv6reg ipv6的正则
const Ipv6Reg = `((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))`

// DNSConfig DNS配置
type DNSConfig struct {
	Name   string
	ID     string
	Secret string
}

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
	DNS DNSConfig
}

func (conf *Config) InitConfigFromFile() error {
	// 从文件中读取配置
	configFilePath := util.GetConfigFromFile()
	_, err := os.Stat(configFilePath)
	if err != nil {
		log.Println("config.yaml 文件不存在")
		return err
	}
	byt, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println("config.yaml 读取失败")
		return err
	}
	log.Println("config.yaml 读取成功")
	// 解析配置
	yaml.Unmarshal(byt, conf)
	// 对byt进行操作，切片解码给到conf
	return nil
}

func (conf *Config) GetIpv4Addr() (result string) {
	// 从配置文件中读取ipv4地址
	if conf.Ipv4.Enable {
		resp, err := http.Get(conf.Ipv4.URL)
		if err != nil {
			// log.Println("获取ipv4地址失败")
			log.Println("ipv4 获取失败", conf.Ipv4.URL)
			return
		}
		defer resp.Body.Close()
		// body, err := ioutill.ReadFile(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("ipv4 结果读取失败", conf.Ipv4.URL)
			return
		}
		comp := regexp.MustCompile(Ipv4Reg)
		result = comp.FindString(string(body))
	}
	return
}

func (conf *Config) GetIpv6Addr() (result string) {
	// 从配置文件中读取ipv6地址
	if conf.Ipv6.Enable {
		resp, err := http.Get(conf.Ipv4.URL)
		if err != nil {
			log.Println("ipv6 获取失败", conf.Ipv6.URL)
		}
		defer resp.Body.Close()
		// body, err := ioutill.ReadFile(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("ipv6结果读取失败", conf.Ipv6.URL)
			return
		}
		comp := regexp.MustCompile(Ipv6Reg)
		result = comp.FindString(string(body))
	}
	return

}
