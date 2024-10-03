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
    - Linux环境推荐在 `--net=host` 下使用
   
```
docker run -d \
    --name myddns \
    --restart=always \
    --net=host \
    alongnice/myddns
```
- [可选] 使用IPV6后，建议设置登录用户名和密码

- 在docker主机上打开[http://127.0.0.1:12138](http://127.0.0.1:12138)，修改你的配置，成功
![avatar](myddns.png)