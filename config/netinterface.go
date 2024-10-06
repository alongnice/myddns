package config

import (
	"fmt"
	"net"
)

// NetInterface 本机网络
type NetInterface struct {
	Name    string
	Address []string
}

// GetNetInterface 获得网卡地址 返回ipv4和ipv6
func GetNetInterface() (ipv4NetInterfaces []NetInterface, ipv6NetInterfaces []NetInterface, err error) {
	//获取所有网络接口的列表
	allNetInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return ipv4NetInterfaces, ipv6NetInterfaces, err
	}

	//遍历所有网络接口列表
	for i := 0; i < len(allNetInterfaces); i++ {
		// Check if the interface is up
		if (allNetInterfaces[i].Flags & net.FlagUp) != 0 {
			//获取接口的地址列表
			addrs, _ := allNetInterfaces[i].Addrs()
			ipv4 := []string{}
			ipv6 := []string{}

			//遍历地址列表
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
					//获取蒙版前缀尺寸和蒙版尺寸
					maskPrefixSize, maskSize := ipnet.Mask.Size()
					// 128位的掩码为IPV6
					if maskSize == 128 && maskPrefixSize != 128 {
						ipv6 = append(ipv6, ipnet.IP.String())
					}
					// 32位的掩码为IPV4
					if maskSize == 32 {
						ipv4 = append(ipv4, ipnet.IP.String())
					}
				}
			}

			//如果有IPv4地址，请将其附加到IPv4网络接口列表
			if len(ipv4) > 0 {
				ipv4NetInterfaces = append(
					ipv4NetInterfaces,
					NetInterface{
						Name:    allNetInterfaces[i].Name,
						Address: ipv4,
					},
				)
			}
			if len(ipv6) > 0 {
				ipv6NetInterfaces = append(
					ipv6NetInterfaces,
					NetInterface{
						Name:    allNetInterfaces[i].Name,
						Address: ipv6,
					},
				)
			}

		}
	}

	return ipv4NetInterfaces, ipv6NetInterfaces, nil
}
