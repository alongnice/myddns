package main

import (
	"embed"
	"fmt"
	"myddns/config"
	"myddns/dns"
	"myddns/util"
	"myddns/web"
	"os"
	"path/filepath"
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

// 配置文件路径
var configFilePath = flag.String("c", "util.GetConfigFromFileDefault()", "自定义配置文件路径")

//go:embed static
var staticEmbededFiles embed.FS

//go:embed favicon.ico
var faviconEmbededFile embed.FS

func main() {
	flag.Parse()
	if _, err := net.ResolveTCPAddr("tcp", *listen); err != nil {
		log.Fatalf("解析监听地址异常，%s", err)
	}
	if *configFilePath != "" {
		absPath, _ := filepath.Abs(*configFilePath)
		os.Setenv(util.ConfigFilePathENV, absPath)
	}
	switch *serviceType {
	case "install":
		installService()
	case "uninstall":
		uninstallService()
	default:
		if util.IsRunInDocker() {
			run(5 * time.Millisecond)
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
					log.Println("可使用 ./myddns -s install 安装服务运行")
				}
				run(5 * time.Millisecond)
			}
		}
	}
}

func run(firstDelay time.Duration) {
	// 启动静态文件服务
	http.Handle("/static/", http.FileServer(http.FS(staticEmbededFiles)))
	http.Handle("/favicon.ico", http.FileServer(http.FS(faviconEmbededFile)))

	http.HandleFunc("/", web.BasicAuth(web.Writing))
	http.HandleFunc("/save", web.BasicAuth(web.Save))
	http.HandleFunc("/logs", web.BasicAuth(web.Logs))
	http.HandleFunc("/clearLog", web.BasicAuth(web.ClearLog))
	http.HandleFunc("/ipv4NetInterface", web.BasicAuth(web.IPv4NetInterface))
	http.HandleFunc("/ipv6NetInterface", web.BasicAuth(web.IPv6NetInterface))
	http.HandleFunc("/webhookTest", web.BasicAuth(web.WebhookTest))

	log.Println("监听", *listen, "...")
	log.Println("http://127.0.0.1:12138")

	// 没有配置,自动打开浏览器
	autoOpenExplorer()

	// 定时运行
	go dns.RunTimer(firstDelay, time.Duration(*every)*time.Second)
	err := http.ListenAndServe(*listen, nil)

	if err != nil {
		log.Println("启动时端口发生异常,检查端口占用情况", err)
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
	options := make(service.KeyValue)
	if service.ChosenSystem().String() == "unix-systemv" {
		options["sysvScript"] = sysvScript
		options["UserService"] = false
	} else if service.ChosenSystem().String() == "linux-upstart" ||
		service.ChosenSystem().String() == "linux-openrc" {
	} else {
		options["UserService"] = true
	}

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
	s.Stop()
	if err := s.Uninstall(); err == nil {
		log.Println("myddns 服务卸载成功!")
	} else {
		log.Printf("myddns 服务卸载失败, ERR: %s\n", err)
	}
}

// 安装服务
func installService() {
	s := getService()
	status, err := s.Status()
	if err != nil && status == service.StatusUnknown {
		if err = s.Install(); err == nil {
			s.Start()
			log.Println("安装 myddns 服务成功! 通过浏览器配置")
			if service.ChosenSystem().String() == "Unix-systemv" {
				log.Println("不能访问，则重启")
			}
			return
		}

		if status != service.StatusUnknown {
			log.Println("myddns 服务已安装, 无需二次安装")
		}
	}

	log.Printf("安装 myddns 服务失败, ERR: %s\n", err)
	switch s.Platform() {
	case "windows-service":
		log.Println("请以管理员身份运行cmd并确保使用如下命令: .\\myddns.exe -s install")
	default:
		log.Println("请确保使用如下命令: ./myddns -s install")
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

// 系统环境脚本
const sysvScript = `#!/bin/sh
# For RedHat and cousins:
# chkconfig: - 99 01
# description: {{.Description}}
# processname: {{.Path}}
### BEGIN INIT INFO
# Provides:          {{.Path}}
# Required-Start:	$local_fs $network $named $time
# Required-Stop:	$local_fs $network $named
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: {{.DisplayName}}
# Description:       {{.Description}}
### END INIT INFO
cmd="{{.Path}}{{range .Arguments}} {{.|cmd}}{{end}}"
name=$(basename $(readlink -f $0))
pid_file="/var/run/$name.pid"
stdout_log="/var/log/$name.log"
stderr_log="/var/log/$name.err"
[ -e /etc/sysconfig/$name ] && . /etc/sysconfig/$name
get_pid() {
    cat "$pid_file"
}
is_running() {
    [ -f "$pid_file" ] && cat /proc/$(get_pid)/stat > /dev/null 2>&1
}
case "$1" in
    start)
        if is_running; then
            echo "Already started"
        else
            echo "Starting $name"
            {{if .WorkingDirectory}}cd '{{.WorkingDirectory}}'{{end}}
            $cmd >> "$stdout_log" 2>> "$stderr_log" &
            echo $! > "$pid_file"
            if ! is_running; then
                echo "Unable to start, see $stdout_log and $stderr_log"
                exit 1
            fi
        fi
    ;;
    stop)
        if is_running; then
            echo -n "Stopping $name.."
            kill $(get_pid)
            for i in $(seq 1 10)
            do
                if ! is_running; then
                    break
                fi
                echo -n "."
                sleep 1
            done
            echo
            if is_running; then
                echo "Not stopped; may still be shutting down or shutdown may have failed"
                exit 1
            else
                echo "Stopped"
                if [ -f "$pid_file" ]; then
                    rm "$pid_file"
                fi
            fi
        else
            echo "Not running"
        fi
    ;;
    restart)
        $0 stop
        if is_running; then
            echo "Unable to stop, will not attempt to start"
            exit 1
        fi
        $0 start
    ;;
    status)
        if is_running; then
            echo "Running"
        else
            echo "Stopped"
            exit 1
        fi
    ;;
    *)
		$0 start
    ;;
esac
exit 0
`
