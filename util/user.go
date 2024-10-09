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
		log.Println("获取路径失败!")
		return "../.ddns_go_config.yaml"
	}
	return dir + string(os.PathSeparator) + ".ddns_go_config.yaml"
}
