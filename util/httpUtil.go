package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// 获得并简单解析 http 接口返回结果
func GetHTTPResponse(resp *http.Response, url string, err error, result interface{}) error {
	if err != nil {
		log.Println("接口请求 失败： ", err)
	} else if resp.StatusCode != 200 {
		log.Println("接口请求 失败： ", resp.Status, "statuscode is ", resp.StatusCode)
		return errors.New("请求返回结果不是200")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("接口请求 失败： ", err, "URL", url)
		}
		err = json.Unmarshal(body, &result)

		if err != nil {
			log.Println("接口请求 ： ", err, "URL", url, "json解析失败", err)
		}
	}

	return err
}
