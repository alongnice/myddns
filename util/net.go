package util

import (
	"net"
	"strings"
)

// 判断是否为私有地址
// https://en.wikipedia.org/wiki/Private_network
func IsPrivateNetwork(remoteAddr string) bool {
	lastIndex := strings.LastIndex(remoteAddr, ":")
	// 获取地址中 : 的位置
	if lastIndex < 1 {
		return false
	}
	// 如果在开头或者结尾，则不是合法的地址

	remoteAddr = remoteAddr[:lastIndex]
	// 截取 : 之前的内容

	// ipv6
	if strings.HasPrefix(remoteAddr, "[") && strings.HasSuffix(remoteAddr, "]") {
		remoteAddr = remoteAddr[1 : len(remoteAddr)-1]
	}
	// 去除 []

	if ip := net.ParseIP(remoteAddr); ip != nil {
		if ip.IsLoopback() {
			return true
		}
		// 回环地址 127.0.0.1

		_, ipNet192, _ := net.ParseCIDR("192.168.0.0/16")
		if ipNet192.Contains(ip) {
			return true
		}

		_, ipNet172, _ := net.ParseCIDR("172.16.0.0/12")
		if ipNet172.Contains(ip) {
			return true
		}

		_, ipNet10, _ := net.ParseCIDR("10.0.0.0/8")
		if ipNet10.Contains(ip) {
			return true
		}

		_, ipNet100, _ := net.ParseCIDR("100.0.0.0/8")
		if ipNet100.Contains(ip) {
			return true
		}

		_, ipNetFE, _ := net.ParseCIDR("fe80::/10")
		if ipNetFE.Contains(ip) {
			return true
		}

		_, ipNetV6FD, _ := net.ParseCIDR("fd00::/8")
		if ipNetV6FD.Contains(ip) {
			return true
		}

	}

	return false
}
