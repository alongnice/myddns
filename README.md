# myddns
- DDNS 动态域名系统，Dynamic DNS
- 服务端在网络不稳定状态下，ip发生变化
- 本服务目的用于周期发布ip给到DNS
- 复刻自 ddns-go


> go-bindata -pkg web -o web/pages-data.go static/pages/...
> 
> go-bindata -fs -prefix "static/" static/