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
## 发布
```
go-bindata -pkg util -o util/staticPagesData.go static/pages/...
go-bindata -pkg static -o static/js_css_data.go -fs -prefix "static/" static/
```


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