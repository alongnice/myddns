package main

import (
	"embed"
	"myddns/config"
	"myddns/dns"
	"myddns/util"
	"myddns/web"
	"os"

	"flag"
	"net"

	"log"
	"net/http"
	"time"
)

// 监听地址
var listen = flag.String("l", ":12138", "web server listen address")

// 更新频率
var every = flag.Int("f", 300, "dns update frequency in second")

//go:embed static
var staticEmbededFiles embed.FS

//go:embed favicon.ico
var faviconEmbededFile embed.FS

func main() {
	flag.Parse()
	if _, err := net.ResolveTCPAddr("tcp", *listen); err != nil {
		log.Fatalf("解析监听地址异常，%s", err)
	}
	if util.IsRunInDocker() {
		run()
	} else {
		runAsService()
	}
}

func run() {
	// 启动静态文件服务
	http.Handle("/static/", http.FileServer(http.FS(staticEmbededFiles)))
	http.Handle("/favicon.ico", http.FileServer(http.FS(faviconEmbededFile)))

	http.HandleFunc("/", config.BasicAuth(web.Writing))
	http.HandleFunc("/save", config.BasicAuth(web.Save))
	http.HandleFunc("/logs", config.BasicAuth(web.Logs))
	http.HandleFunc("/ipv4NetInterface", config.BasicAuth(web.IPv4NetInterface))
	http.HandleFunc("/ipv6NetInterface", config.BasicAuth(web.IPv6NetInterface))

	log.Println("监听", *listen, "...")

	// 定时运行
	go dns.RunTimer(time.Duration(*every) * time.Second)
	err := http.ListenAndServe(*listen, nil)

	if err != nil {
		log.Println("启动时端口发生异常,1分钟内自动关闭窗口", err)
		time.Sleep(time.Minute)
		os.Exit(1)
	}
}
