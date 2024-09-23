# myddns
- DDNS 动态域名系统，Dynamic DNS
- 服务端在网络不稳定状态下，ip发生变化
- 本服务目的用于周期发布ip给到DNS
- 复刻自 ddns-go



## 依赖
```

 go get -u github.com/go-bindata/go-bindata/...
 go-bindata -debug -pkg util -o util/staticPagesData.go static/pages/...
 go-bindata -pkg static -o static/js_css_data.go -fs -prefix "static/" static/
 ```