package util

import (
	"log"
	"os"
	"os/user"
)

func GetConfigFromFile() string {
	u, err := user.Current()
	if err != nil {
		log.Println("获取用户信息失败")
		return "../.myddns_conf.yaml"
	}
	return u.HomeDir + string(os.PathSeparator) + ".myddns_conf.yaml"
}
