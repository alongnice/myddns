package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

// 获得并简单解析 http 接口返回结果
func GetHTTPResponse(resp *http.Response, url string, err error, result interface{}) error {
	if err != nil {
		log.Println("接口请求 失败： ", err)
	// } else if resp.StatusCode != 200 {
	// 	defer resp.Body.Close()
	// 	log.Println("接口请求 失败： ", resp.Status, "statuscode is ", resp.StatusCode)
	// 	return errors.New("请求返回结果不是200")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("接口请求 失败： ", err, "URL", url)
		}
		if resp.StatusCode != 200 {
			log.Printf("请求接口 %s 失败! %s \n",url, string(body))
			err = fmt.Errorf("请求接口 %s 失败! %s \n",url, string(body))
		}

		err = json.Unmarshal(body, &result)

		if err != nil {
			log.Println("接口请求 ： ", err, "URL", url, "json解析失败", err)
		}
	}

	return err
}
