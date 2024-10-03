# myddns
- DDNS 动态域名系统，Dynamic DNS
- 服务端在网络不稳定状态下，ip发生变化
- 本服务目的用于周期发布用户服务的 公网IP 给到DNS服务器
- 支持多级域名
- 支持跨平台跨架构运行，x86,arm,  *nix, win, mac
- 复刻自 ddns-go
- /logs 子域名 页面快速查看最近50条日志

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
    -p 127.0.0.1:12138:12138
    alongnice/myddns
```

- 在docker主机上打开[http://127.0.0.1:12138](http://127.0.0.1:12138)，修改你的配置，成功
![avatar](myddns.png)