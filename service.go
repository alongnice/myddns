package main

import (
	"fmt"
	"log"
	"myddns/config"
	"myddns/util"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	// 异步执行
	go p.run()
	return nil
}
func (p *program) run() {
	run()
}
func (p *program) Stop(s service.Service) error {
	return nil
}

// 以服务方式运行
func runAsService() {
	svcConfig := &service.Config{
		Name:        "myddns",
		DisplayName: "myddns",
		Description: "简单好用的动态域名解析工具。自动更新域名解析到公网IP(支持阿里云、腾讯云dnspod、Cloudflare、华为云)",
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
			if err = s.Uninstall(); err == nil {
				log.Println("myddns 服务卸载成功!")
			} else {
				log.Printf("ddns-go 服务卸载失败, ERR: %s\n", err)
				switch s.Platform() {
				case "windows-service":
					log.Println("请确保使用如下命令: .\\ddns-go uninstall")
				default:
					log.Println("请确保使用如下命令: sudo ./ddns-go uninstall")
				}
			}
			return
		}
	}

	status, err := s.Status()
	if err != nil && status == service.StatusUnknown {
		// 服务未知，创建服务
		if err = s.Install(); err == nil {
			s.Start()
			openExplorer()
			log.Println("安装 myddns 服务成功! 程序会一直运行, 包括重启后。")
			log.Println("如需卸载 myddns, 使用 sudo ./myddns uninstall")
			log.Println("请在浏览器中进行配置。1分钟后自动关闭DOS窗口!")
			time.Sleep(time.Minute)
			return
		}

		log.Printf("安装 ddns-go 服务失败, ERR: %s\n", err)
		switch s.Platform() {
		case "windows-service":
			log.Println("请确保使用如下命令: .\\ddns-go")
		default:
			log.Println("请确保使用如下命令: sudo ./ddns-go")
		}
	}

	// 正常启动
	s.Run()

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
