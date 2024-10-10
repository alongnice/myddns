package util

const MaxTimes = 5

// IpCache 上次IP缓存
type IpCache struct {
	Addr         string // 缓存地址
	Times        int    // 剩余次数
	ForceCompare bool   // 是否强制比对
}

var Ipv4Cache *IpCache = &IpCache{}
var Ipv6Cache *IpCache = &IpCache{}

func (ca *IpCache) Check(newAddr string) bool {
	if newAddr == "" {
		return true
	}
	// 地址改变 或 达到剩余次数 或 强制比对
	if ca.Addr != newAddr ||
		ca.Times == MaxTimes ||
		ca.ForceCompare {
		ca.Addr = newAddr
		ca.Times = 0
		ca.ForceCompare = false
		return true
	}
	ca.Addr = newAddr
	ca.Times = ca.Times + 1
	return false
}
