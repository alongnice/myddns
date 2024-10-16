package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"myddns/util"
	"os"
	"regexp"
	"strings"
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
		Enable       bool
		GetType      string
		URL          string
		NetInterface string
		Domains      []string
	}
	Ipv6 struct {
		Enable       bool
		GetType      string
		URL          string
		NetInterface string
		Domains      []string
	}
	DNS DNSConfig
	User
	Webhook
	// 禁止从公网访问
	NotAllowWanAccess bool
	TTL               string
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
	cache.Lock.Lock()
	defer cache.Lock.Unlock()

	if cache.ConfigSingle != nil {
		return *cache.ConfigSingle, cache.Err
	}

	// init config
	cache.ConfigSingle = &Config{}

	configFilePath := util.GetConfigFromFile()
	_, err = os.Stat(configFilePath)
	if err != nil {
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	byt, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Println("config.yaml读取失败")
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	err = yaml.Unmarshal(byt, cache.ConfigSingle)
	if err != nil {
		log.Println("反序列化配置文件失败", err)
		cache.Err = err
		return *cache.ConfigSingle, err
	}
	// remove err
	cache.Err = nil
	return *cache.ConfigSingle, err
}

// 保存配置
func (conf *Config) SaveConfig() (err error) {
	cache.Lock.Lock()
	defer cache.Lock.Unlock()

	byt, err := yaml.Marshal(conf)
	if err != nil {
		log.Println(err)
		return err
	}

	configFilePath := util.GetConfigFromFile()
	err = ioutil.WriteFile(configFilePath, byt, 0600)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("配置文件已保存在: %s\n", configFilePath)

	// 清空配置缓存
	cache.ConfigSingle = nil

	return
}

func (conf *Config) GetIpv4Addr() (result string) {
	// 从配置文件中读取ipv4地址
	// 判断从哪里获取ip
	if conf.Ipv4.GetType == "netInterface" {
		// 从网卡获取IP
		ipv4, _, err := GetNetInterface()
		if err != nil {
			log.Println("获取ipv4地址失败")
			return
		}

		for _, netInterface := range ipv4 {
			if netInterface.Name == conf.Ipv4.NetInterface {
				return netInterface.Address[0]
			}
		}
		log.Println("从网卡中获取IPv4失败,网卡名:", conf.Ipv4.NetInterface)
		return
	}

	client := util.CreateHTTPClient()
	// 创建一个http.Client的对象
	// 设置超时时间
	// 禁用keep-DisableKeepAlives
	// 发送http请求
	urls := strings.Split(conf.Ipv4.URL, ",") // 多个地址 改用 , 分割
	for _, url := range urls {
		url = strings.TrimSpace(url) // 去除空格
		resp, err := client.Get(url)
		if err != nil {
			log.Println(fmt.Sprintf("连接失败! <a target='blank' href='%s'>点击查看接口能否返回IPv4地址</a>,", url))
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("读取" + url + "失败")
			continue
		}
		comp := regexp.MustCompile(Ipv4Reg)
		result = comp.FindString(string(body))
		if result != "" {
			return
		} else {
			log.Printf("从%s获取IP失败;错误码 %s\n", url, result)
		}
	}
	return
}

func (conf *Config) GetIpv6Addr() (result string) {
	// 从配置文件中读取ipv6地址
	// 判断从哪里获取ip
	if conf.Ipv4.GetType == "netInterface" {
		// 从网卡获取IP
		ipv4, _, err := GetNetInterface()
		if err != nil {
			log.Println("获取ipv6地址失败")
			return
		}

		for _, netInterface := range ipv4 {
			if netInterface.Name == conf.Ipv4.NetInterface {
				return netInterface.Address[0]
			}
		}
		log.Println("从网卡中获取IPv6失败,网卡名:", conf.Ipv6.NetInterface)
		return
	}

	client := util.CreateHTTPClient()
	urls := strings.Split(conf.Ipv6.URL, ",")
	for _, url := range urls {
		url = strings.TrimSpace(url)
		resp, err := client.Get(url)
		if err != nil {
			log.Println(fmt.Sprintf("连接失败! <a target='blank' href='%s'>点击查看接口能否返回IPv6地址</a>, 官方说明:<a target='blank' href='%s'>点击访问</a> ", url, "https://github.com/alongnice/myddns"))
			continue
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("读取IPv6结果失败! 接口: ", url)
			continue
		}
		comp := regexp.MustCompile(Ipv6Reg)
		result = comp.FindString(string(body))
		if result != "" {
			return
		} else {
			log.Printf("获取IPv6结果失败! 接口: %s ,返回值: %s\n", url, result)
		}
	}

	return

}
