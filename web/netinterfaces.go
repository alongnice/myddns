package web

import (
	"encoding/json"
	"myddns/config"
	"net/http"
)

func IPv4NetInterface(writer http.ResponseWriter, request *http.Request) {
	ipv4, _, err := config.GetNetInterface()
	if len(ipv4) > 0 && err == nil {
		byt, err := json.Marshal(ipv4)
		if err == nil {
			writer.Write(byt)
			return
		}
	}
}

func IPv6NetInterface(writer http.ResponseWriter, request *http.Request) {
	ipv6, _, err := config.GetNetInterface()
	if len(ipv6) > 0 && err == nil {
		byt, err := json.Marshal(ipv6)
		if err == nil {
			writer.Write(byt)
			return
		}
	}
}
