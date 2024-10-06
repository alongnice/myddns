package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"myddns/util"
	"net/http"
	"os"
	"regexp"
	"sync"

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
	User
	Webhook
}

// 配置缓存
type cacheType struct {
	ConfigSingle *Config
	Err          error
	Lock         sync.Mutex
}

var cache = &cacheType{}

// 获得配置
// func (conf *Config) InitConfigFromFile() error {
func GetConfigCache() (conf Config, err error) {
	if cache.ConfigSingle != nil {
		return *cache.ConfigSingle, cache.Err
	}
	cache.Lock.Lock()
	defer cache.Lock.Unlock()

	cache.ConfigSingle = &Config{}

	// 从文件中读取配置
	configFilePath := util.GetConfigFromFile()
	_, err = os.Stat(configFilePath)
	if err != nil {
		log.Println("config.yaml 文件不存在,请输入配置文件")
		cache.Err = err
		return *cache.ConfigSingle, err
	}
	byt, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Println("config.yaml 读取失败")
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	err = yaml.Unmarshal(byt, cache.ConfigSingle)
	if err != nil {
		log.Println("反序列化配置文件,失败", err)
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	cache.Err = nil
	// 对byt进行操作，切片解码给到conf
	return *cache.ConfigSingle, err
}

// 保存配置
func (conf *Config) SaveConfig() (err error) {
	byt, err := yaml.Marshal(conf)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(util.GetConfigFromFile(), byt, 0600)
	if err != nil {
		log.Println(err)
		return
	}
	// 清空配置缓存
	cache.ConfigSingle = nil

	return
}

func (conf *Config) GetIpv4Addr() (result string) {
	// 从配置文件中读取ipv4地址
	if conf.Ipv4.Enable {
		resp, err := http.Get(conf.Ipv4.URL)
		if err != nil {
			// log.Println("获取ipv4地址失败")
			log.Println(fmt.Sprintf("未能获得IPV4地址! <a target='blank' href='%s'>点击查看接口能否返回IPV4地址</a>,", conf.Ipv4.URL))
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
			log.Println(fmt.Sprintf("未能获得IPV6地址! <a target='blank' href='%s'>点击查看接口能否返回IPV6地址</a>, 官方说明:<a target='blank' href='%s'>点击访问</a> ", conf.Ipv6.URL, "https://github.com/jeessy2/ddns-go#使用ipv6"))
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
