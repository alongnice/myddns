package main

import (
	"myddns/util"
	"myddns/web"

	"log"
	"net/http"
	"time"
)

const port = "12138"

func main() {
	// 启动静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", web.Writing)
	http.HandleFunc("/save", web.Save)

	// 新建协程,打开浏览器
	go util.OpenExplore("http://127.0.0.1:" + port)
	log.Println(port, " 端口启动", "...")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("启动时端口发生异常,1分钟内自动关闭窗口", err)
		time.Sleep(time.Minute)
	}
}
