package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"net/url"
)

// https://github.com/rosbit/aliyun-sign/blob/master/aliyun-sign.go

var (
	signMethodMap = map[string]func() hash.Hash{
		"HMAC-SHA1":   sha1.New,
		"HMAC-SHA256": sha256.New,
		"HMAC-MD5":    md5.New,
	}
)

// HmacSign 函数用于生成HMAC签名
func HmacSign(signMethod string, httpMethod string, appKeySecret string, vals url.Values) (signature []byte) {
	// 将appKeySecret和"&"拼接成key
	key := []byte(appKeySecret + "&")

	// 声明一个hash.Hash类型的变量h
	var h hash.Hash
	// 判断signMethod是否在signMethodMap中
	if method, ok := signMethodMap[signMethod]; ok {
		// 如果在，则使用对应的hash方法
		h = hmac.New(method, key)
	} else {
		// 如果不在，则使用sha1.New方法
		h = hmac.New(sha1.New, key)
	}
	// 调用makeDataToSign函数，将httpMethod和vals作为参数传入
	makeDataToSign(h, httpMethod, vals)
	// 返回h.Sum(nil)的值
	return h.Sum(nil)
}

// HmacSignToB64 函数用于对给定的参数进行Hmac签名，并将签名结果进行Base64编码
func HmacSignToB64(signMethod string, httpMethod string, appKeySecret string, vals url.Values) (signature string) {
	// 对给定的参数进行Hmac签名
	return base64.StdEncoding.EncodeToString(HmacSign(signMethod, httpMethod, appKeySecret, vals))
}

type strToEnc struct {
	s string
	e bool
}

// makeDataToSign函数用于生成待签名的数据
func makeDataToSign(w io.Writer, httpMethod string, vals url.Values) {
	// 创建一个通道，用于传递待签名的数据
	in := make(chan *strToEnc)
	// 启动一个goroutine，用于生成待签名的数据
	go func() {
		// 将httpMethod放入通道
		in <- &strToEnc{s: httpMethod}
		// 将"&"放入通道
		in <- &strToEnc{s: "&"}
		// 将"/"放入通道，并设置e为true，表示这是最后一个元素
		in <- &strToEnc{s: "/", e: true}
		// 将"&"放入通道
		in <- &strToEnc{s: "&"}
		// 将vals.Encode()放入通道，并设置e为true，表示这是最后一个元素
		in <- &strToEnc{s: vals.Encode(), e: true}
		// 关闭通道
		close(in)
	}()

	// 调用specialUrlEncode函数，将待签名的数据写入w
	specialUrlEncode(in, w)
}

var (
	encTilde = "%7E"         // '~' -> "%7E"
	encBlank = []byte("%20") // ' ' -> "%20"
	tilde    = []byte("~")
)

// specialUrlEncode函数用于对输入的字符串进行特殊URL编码
func specialUrlEncode(in <-chan *strToEnc, w io.Writer) {
	for s := range in {
		// 如果字符串不需要编码，则直接写入
		if !s.e {
			io.WriteString(w, s.s)
			continue
		}

		// 获取字符串的长度
		l := len(s.s)
		// 遍历字符串
		for i := 0; i < l; {
			ch := s.s[i]

			// 判断字符是否为特殊字符
			switch ch {
			case '%':
				// 如果字符为'~'，则替换为"%7E"
				if encTilde == s.s[i:i+3] {
					w.Write(tilde)
					i += 3
					continue
				}
				fallthrough
			case '*', '/', '&', '=':
				// 如果字符为'*'、'/'、'&'、'='，则替换为"%XX"
				fmt.Fprintf(w, "%%%02X", ch)
			case '+':
				// 如果字符为'+'，则替换为"%20"
				w.Write(encBlank)
			default:
				// 其他字符直接写入
				fmt.Fprintf(w, "%c", ch)
			}

			i += 1
		}
	}
}
