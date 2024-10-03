# myddns
- DDNS 动态域名系统，Dynamic DNS
- 服务端在网络不稳定状态下，ip发生变化
- 本服务目的用于周期发布用户服务的 公网IP 给到DNS服务器
- 支持多级域名
- 支持跨平台跨架构运行，x86,arm,  *nix, win, mac
- /logs 子域名 页面快速查看最近50条日志 简化日志审查
- 网页中简单配置,可设置用户名账号密码
- 复刻自 ddns-go

## 可选域名供应商
+ Alidns 阿里云
+ Dnspod 腾讯云
+ cloudflare
+ 华为云

## 依赖
```
 go get -u github.com/go-bindata/go-bindata/...
 go-bindata -debug -pkg util -o util/staticPagesData.go static/pages/...
 go-bindata -pkg static -o static/js_css_data.go -fs -prefix "static/" static/
```
## 发布
```
go-bindata -pkg util -o util/staticPagesData.go static/pages/...
go-bindata -pkg static -o static/js_css_data.go -fs -prefix "static/" static/
```

## 普通环境使用
- 下载[https://github.com/alongnice/myddns/releases](https://github.com/alongnice/myddns/releases)
- 运行，程序将自行打开浏览器，访问 [http://127.0.0.1:12138](http://127.0.0.1:12138)完成配置修改




## Docker使用
```
docker run -d \
    --name myddns \ 
    --restart=always \
    -p 12138:12138
    alongnice/myddns
```

## 使用 IPV6
    - 前提: 你的环境需要支持 IPV6
    - Windows/Mac 系统推荐在 `系统中使用`, windows/mac 桌面的docker不支持主机网络 `--net=host`
    - Linux的x86或arm架构，如服务器、群晖(网络中勾选`使用与docker相同的网络`)、xx盒子等等，推荐使用`--net=host`模式
    - 
    - 群晖：使用docker。
    - 1. 注册表中搜索`ddns-go`并下载。 
    - 2. 映像 -> 选择`jeessy/ddns-go` -> 启动 -> 高级设置 -> 网络中勾选`使用与 Docker Host 相同的网络`)
   
```
docker run -d \
    --name myddns \
    --restart=always \
    --net=host \
    alongnice/myddns
```
- 虚拟机存在可能能获取IPV6,但是无法正常访问
- [可选] 使用IPV6后，建议设置登录用户名和密码

## Webhook
- 支持webhook, 当IP变化, 会回调填写的URL
- 支撑变量
  
| 变量名 | 说明 |
| ----- | ---- |
| #{ipv4New}| 新的ipv4地址 |
| #{ipv4Old}| 旧的ipv4地址 |
| #{ipv6New}| 新的ipv6地址 |
| #{ipv6Old}| 旧的ipv6地址 |
| #{ipv4Domains}|Ipv4的域名,多个域名用逗号分隔 |
| #{ipv6Domains}|Ipv6的域名,多个域名用逗号分隔 |

- RequestBody 为空 GET 请求, 不为空 POST 请求
- 示例
- URL:  `https://sc.ftqq.com/[SCKEY].send?text=主人IP变了#{ipv4New}`
- RequestBody:  `{"text":"你的IPv4已变为#{ipv4New}","desp":"域名有#{ipv4Domains}"}}`

---

- 在docker主机上打开[http://127.0.0.1:12138](http://127.0.0.1:12138)，修改你的配置，成功
![avatar](myddns.png)