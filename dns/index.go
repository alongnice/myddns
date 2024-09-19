package dns

import "myddns/config"

type DNS interface {
	AddRecord(conf *config.Config) (ipv4 bool, ipv6 bool)
}
