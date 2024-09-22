package util

import (
	"io/ioutil"
	"log"
	// "os"
	"strings"
	"fmt"
)

// GetStaticResourcePath 获得静态资源文件路径
func GetStaticResourcePath(orgPath string) (temPath string, err error) {
	data, err := Asset(orgPath)
	if err != nil {
		log.Println("Asset was not found.")
		return "", err
	}
	// tempFile := os.TempDir() + strings.ReplaceAll(orgPath, "/", "/")
	fmt.Println(orgPath)
	tempFile := strings.ReplaceAll(orgPath, "/", "/")
	// ioutil.WriteFile(tempFile, data, 0600)
	ioutil.WriteFile(temPath, data, 0600)
	return tempFile, nil
}
