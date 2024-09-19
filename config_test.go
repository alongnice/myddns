package main

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConfig(t *testing.T) {
	conf := &Config{}
	byt, err := os.ReadFile("config.yaml")
	if err != nil {
		t.Error(err)
	}

	yaml.Unmarshal(byt, conf)
	if "alidns" != conf.DNS.Name {
		t.Error("dns name 错误，暂时只支持 alidns")
	}
}
