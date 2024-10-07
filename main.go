package main

import (
	"myddns/config"
	"myddns/dns"
	static "myddns/static"
	"myddns/util"
	"myddns/web"

	"flag"
	"fmt"
	"net"
	"strconv"

	"log"
	"net/http"
	"time"
)

func main() {
	listen := flag.String("l", ":12138", "web server listen address")
	every := flag.String("f", "300", "dns update frequency in second")
	flag.Parse()

	// 启动静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static.AssetFile())))
	http.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(static.AssetFile())))

	http.HandleFunc("/", config.BasicAuth(web.Writing))
	http.HandleFunc("/save", config.BasicAuth(web.Save))
	http.HandleFunc("/logs", config.BasicAuth(web.Logs))
	http.HandleFunc("/ipv4NetInterface", config.BasicAuth(web.IPv4NetInterface))
	http.HandleFunc("/ipv6NetInterface", config.BasicAuth(web.IPv6NetInterface))

	addr, err := net.ResolveTCPAddr("tcp", *listen)
	if err != nil {
		log.Println("解析监听地址异常", err)
	}
	url := ""
	if addr.IP.IsGlobalUnicast() {
		url = fmt.Sprintf("http://%s", addr.String())
	} else if addr.IP.To4() != nil || addr.IP == nil || addr.IP.Equal(net.ParseIP("::")) {
		url = fmt.Sprintf("http://127.0.0.1:%d", addr.Port)
	} else {
		url = fmt.Sprintf("http://[::1]:%d", addr.Port)
	}

	// 新建协程,打开浏览器
	_, err = config.GetConfigCache()
	if err != nil {
		go util.OpenExplore(url)
	}
	log.Println("监听", *listen, "...")

	// 定时运行
	delay, err := strconv.ParseUint(*every, 10, 64)
	if err != nil {
		delay = 300
	}
	go dns.RunTimer(time.Duration(delay) * time.Second)
	err = http.ListenAndServe(*listen, nil)

	if err != nil {
		log.Println("启动时端口发生异常,1分钟内自动关闭窗口", err)
		time.Sleep(time.Minute)
	}
}
