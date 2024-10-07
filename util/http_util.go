package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
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
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		log.Printf("请求接口 %s 失败! %s \n",url, string(body))
		err = fmt.Errorf("请求接口 %s 失败! %s \n",url, string(body))
	}
	

	return body, err
}
