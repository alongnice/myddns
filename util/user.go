package util

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

const ConfigFilePathENV = "MYDDNS_CONFIG_FILE_PATH"

func GetConfigFromFile() string {
	configFilePath := os.Getenv(ConfigFilePathENV)
	if configFilePath != "" {
		return configFilePath
	}
	return GetConfigFromFileDefault()
}

func GetConfigFromFileDefault() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Println("获取用户信息失败")
		return "../.myddns_conf.yaml"
	}
	return dir + string(os.PathSeparator) + ".myddns_conf.yaml"
}
