package main

import (
	"myddns/dns"
	"myddns/static"
	"myddns/util"
	"myddns/web"

	"log"
	"net/http"
	"time"
)

const port = "12138"

func main() {
	// 启动静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static.AssetFile())))
	http.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(static.AssetFile())))

	http.HandleFunc("/", web.Writing)
	http.HandleFunc("/save", web.Save)
	// 添加日志模块
	http.HandleFunc("/logs", web.Logs)

	// 新建协程,打开浏览器
	go util.OpenExplore("http://127.0.0.1:" + port)
	log.Println(port, " 端口启动", "...")

	// 定时运行
	go dns.RunTimer()

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("启动时端口发生异常,1分钟内自动关闭窗口", err)
		time.Sleep(time.Minute)
	}
}
