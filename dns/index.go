package dns

type DNS interface {
	addRecord() (ipv4 bool, ipv6 bool)
}
