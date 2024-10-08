package web

import (
	"bytes"
	"encoding/base64"
	"log"
	"myddns/config"
	"myddns/util"
	"net/http"
	"strings"
	"time"
)

// viewFunc
type ViewFunc func(http.ResponseWriter, *http.Request)

// 登录重试次数
var loginRetryCount struct {
	FailTimes int
}

var ld = &loginRetryCount

// BasicAuth 登录验证
func BasicAuth(f ViewFunc) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 账号或密码为空跳过
		conf, _ := config.GetConfigCache()

		// 禁止公网访问
		if conf.NotAllowWanAccess {
			if !util.IsPrivateNetwork(r.RemoteAddr) || !util.IsPrivateNetwork(r.Host) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		// 禁止从公网访问
		if conf.NotAllowWanAccess {
			if !util.IsPrivateNetwork(r.RemoteAddr) || !util.IsPrivateNetwork(r.Host) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		if conf.Username == "" && conf.Password == "" {
			// 执行被装饰的函数
			f(w, r)
			return
		}

		// 认证操作
		BasicAuthPrefix := "Basic"

		// 获取请求头
		auth := r.Header.Get("Authorization")
		// 如果是http 登录认证
		if strings.HasPrefix(auth, BasicAuthPrefix) {
			// 解码认证信息
			payload, err := base64.StdEncoding.DecodeString(
				auth[len(BasicAuthPrefix):],
			)

			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 &&
					bytes.Equal(pair[0], []byte(conf.Username)) &&
					bytes.Equal(pair[1], []byte(conf.Password)) {
					// 执行被装饰的函数
					ld.FailTimes = 0
					f(w, r)
					return
				}
			}
			ld.FailTimes++
			if ld.FailTimes > 5 {
				log.Printf("%s:登录失败次数过多", r.RemoteAddr)
				time.Sleep(time.Second * 60)
				ld.FailTimes = 0
				return
			}
			log.Printf("%s:认证失败", r.RemoteAddr)
		}
		// 认证失败
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		// 401 状态码(无权限)
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("%s:请求登录", r.RemoteAddr)
	}
}
