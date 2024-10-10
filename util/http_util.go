package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// 获得并简单解析 http 接口返回json结果
func GetHTTPResponse(resp *http.Response, url string, err error, result interface{}) error {
	body, err := GetHTTPResponseOrg(resp, url, err)
	if err != nil {
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Println("接口请求 ： ", err, "URL", url, "json解析失败", err)
		}
	}
	return err
}

// 获得并简单解析 http 接口返回byte结果
func GetHTTPResponseOrg(resp *http.Response, url string, err error) ([]byte, error) {
	if err != nil {
		log.Println("接口请求", url, " 失败： ", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("接口请求 失败： ", err, "URL", url)
	}
	if resp.StatusCode >= 300 {
		errMsg := fmt.Sprintf("请求接口 %s 失败! %s \n 返回状态码", url, string(body), resp.StatusCode)
		log.Println(errMsg)
		err = fmt.Errorf(errMsg)
	}

	return body, err
}

func CreateHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			// 传输层(TCP?)
			Dial: (&net.Dialer{
				Timeout: 1 * time.Second,
				// 连接超时时间5s
				KeepAlive: 30 * time.Second,
				// 连接保持时间30s
			}).Dial,
			IdleConnTimeout: 10 * time.Second,
			// 空闲链接超时时间
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
	// 与linux平台的posix实现更加丰富 管控参数更加自由
}
