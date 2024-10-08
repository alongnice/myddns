package util

import (
	"log"
	"os"
	"os/user"
)

const ConfigFilePathENV = "MYDDNS_CONFIG_FILE_PATH"

func GetConfigFromFile() string {
	configFilePath := os.Getenv(ConfigFilePathENV)
	if configFilePath != "" {
		return configFilePath
	}
	u, err := user.Current()
	if err != nil {
		log.Println("获取用户信息失败")
		return "../.myddns_conf.yaml"
	}
	return u.HomeDir + string(os.PathSeparator) + ".myddns_conf.yaml"
}
