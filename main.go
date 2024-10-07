package main

import (
	"embed"
	"fmt"
	"myddns/config"
	"myddns/dns"
	"myddns/util"
	"myddns/web"
	"os"
	"strconv"

	"flag"
	"net"

	"log"
	"net/http"
	"time"

	"github.com/kardianos/service"
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

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	run()
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

// 以服务方式运行
func runAsService() {
	svcConfig := &service.Config{
		Name:        "myddns",
		DisplayName: "myddns",
		Description: "简单好用的DDNS。自动更新域名解析到公网IP(支持阿里云、腾讯云dnspod、Cloudflare、华为云)",
		Arguments:   []string{"-l", *listen, "-f", strconv.Itoa(*every)},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// 处理卸载
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "uninstall":
			s.Stop()
			if err = s.Uninstall(); err == nil {
				log.Println("myddns 服务卸载成功!")
			} else {
				log.Printf("myddns 服务卸载失败, ERR: %s\n", err)
				switch s.Platform() {
				case "windows-service":
					log.Println("请以管理员身份运行cmd并确保使用如下命令: .\\myddns.exe uninstall")
				default:
					log.Println("请确保使用如下命令: sudo ./myddns uninstall")
				}
			}
			return
		}
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	status, err := s.Status()
	if err != nil && status == service.StatusUnknown {
		// 服务未知，创建服务
		if err = s.Install(); err == nil {
			s.Start()
			openExplorer()
			log.Println("安装 myddns 服务成功! 程序会一直运行, 包括重启后。")
			switch s.Platform() {
			case "windows-service":
				log.Println("如需卸载 myddns, 请以管理员身份运行cmd并确保使用如下命令: .\\myddns.exe uninstall")
			default:
				log.Println("如需卸载 myddns, 请确保使用如下命令: sudo ./myddns uninstall")
			}
			log.Println("请在浏览器中进行配置。1分钟后自动关闭DOS窗口!")
			time.Sleep(time.Minute)
			return
		}

		log.Printf("安装 myddns 服务失败, ERR: %s\n", err)
		switch s.Platform() {
		case "windows-service":
			log.Println("请以管理员身份运行cmd并确保使用如下命令: .\\myddns.exe")
		default:
			log.Println("请确保使用如下命令: sudo ./myddns")
		}
	}

	// 正常启动
	s.Run()
	if err != nil {
		logger.Error(err)
	}

}

func openExplorer() {
	_, err := config.GetConfigCache()
	// 未找到配置文件&&不是在docker中运行 才打开浏览器
	if err != nil && !util.IsRunInDocker() {
		addr, err := net.ResolveTCPAddr("tcp", *listen)
		if err != nil {
			return
		}
		url := ""
		if addr.IP.IsGlobalUnicast() {
			url = fmt.Sprintf("http://%s", addr.String())
		} else {
			url = fmt.Sprintf("http://127.0.0.1:%d", addr.Port)
		}
		go util.OpenExplore(url)
	}
}
