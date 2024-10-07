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
var listen = flag.String("l", ":12138", "监听地址")

// 更新频率
var every = flag.Int("f", 300, "同步间隔时间(秒)")

// 服务类别
var serviceType = flag.String("s", "", "服务管理, 支持install, uninstall")

//go:embed static
var staticEmbededFiles embed.FS

//go:embed favicon.ico
var faviconEmbededFile embed.FS

func main() {
	flag.Parse()
	if _, err := net.ResolveTCPAddr("tcp", *listen); err != nil {
		log.Fatalf("解析监听地址异常，%s", err)
	}
	switch *serviceType {
	case "install":
		installService()
	case "uninstall":
		uninstallService()
	default:
		if util.IsRunInDocker() {
			run(100 * time.Millisecond)
		} else {
			s := getService()
			status, _ := s.Status()
			if status != service.StatusUnknown {
				// 以服务方式运行
				s.Run()
			} else {
				// 非服务方式运行
				switch s.Platform() {
				case "windows-service":
					log.Println("可使用 .\\myddns.exe -s install 安装服务运行")
				default:
					log.Println("可使用 sudo ./myddns -s install 安装服务运行")
				}
				run(100 * time.Millisecond)
			}
		}
	}
}

func run(firstDelay time.Duration) {
	// 启动静态文件服务
	http.Handle("/static/", http.FileServer(http.FS(staticEmbededFiles)))
	http.Handle("/favicon.ico", http.FileServer(http.FS(faviconEmbededFile)))

	http.HandleFunc("/", config.BasicAuth(web.Writing))
	http.HandleFunc("/save", config.BasicAuth(web.Save))
	http.HandleFunc("/logs", config.BasicAuth(web.Logs))
	http.HandleFunc("/ipv4NetInterface", config.BasicAuth(web.IPv4NetInterface))
	http.HandleFunc("/ipv6NetInterface", config.BasicAuth(web.IPv6NetInterface))
	http.HandleFunc("/webhookTest", config.BasicAuth(web.WebhookTest))

	log.Println("监听", *listen, "...")

	// 没有配置,自动打开浏览器
	autoOpenExplorer()

	// 定时运行
	go dns.RunTimer(firstDelay, time.Duration(*every)*time.Second)
	err := http.ListenAndServe(*listen, nil)

	if err != nil {
		log.Println("启动时端口发生异常,1分钟内自动关闭窗口", err)
		time.Sleep(time.Minute)
		os.Exit(1)
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	log.Printf("以服务方式运行中,配置文件地址: %s\n", util.GetConfigFromFile())
	// 延时运行等待网络
	run(10 * time.Second)
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

// 以服务方式运行
func getService() service.Service {
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
	return s
}

// 卸载服务
func uninstallService() {
	s := getService()
	status, _ := s.Status()
	// 处理卸载
	if status != service.StatusUnknown {
		s.Stop()
		if err := s.Uninstall(); err == nil {
			log.Println("myddns 服务卸载成功!")
		} else {
			log.Printf("myddns 服务卸载失败, ERR: %s\n", err)
		}
	} else {
		log.Printf("myddns 服务未安装")
	}
}

// 安装服务
func installService() {
	s := getService()
	status, err := s.Status()
	if err != nil && status == service.StatusUnknown {
		// 服务未知，创建服务
		if err = s.Install(); err == nil {
			s.Start()
			log.Println("安装 myddns 服务成功! 程序会一直运行, 包括重启后。")
			log.Println("不存在配置文件,请先在浏览器中进行配置。")
			return
		}
		log.Printf("安装 myddns 服务失败, ERR: %s\n", err)

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
		log.Println("请以管理员身份运行cmd并确保使用如下命令: .\\myddns.exe -s install")
	default:
		log.Println("请确保使用如下命令: sudo ./myddns -s install")
	}
	if status != service.StatusUnknown {
		log.Println("myddns 服务已安装, 无需在次安装")
	}
}

// 打开浏览器,理解为重载了 util/open_explore.go 中的实现
func autoOpenExplorer() {
	_, err := config.GetConfigCache()
	// 未找到配置文件
	if err != nil {
		if util.IsRunInDocker() {
			// docker运行, 提示
			fmt.Println("Docker中运行, 请在浏览器中打开 http://docker主机IP:端口 进行配置")
		} else {
			// 主机运行
			addr, err := net.ResolveTCPAddr("tcp", *listen)
			if err != nil {
				return
			}
			url := fmt.Sprintf("http://127.0.0.1:%d", addr.Port)
			if addr.IP.IsGlobalUnicast() {
				url = fmt.Sprintf("http://%s", addr.String())
			}
			go util.OpenExplore(url)
		}
	}
}
