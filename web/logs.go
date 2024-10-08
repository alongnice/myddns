package web

import (
	"io"
	"log"
	"net/http"
	"os"
)

// 内存中的日志
type memLogs struct {
	MaxNum int      // 最大条目
	Logs   []string //日志
}

func (mlog *memLogs) Write(p []byte) (n int, err error) {
	mlog.Logs = append(mlog.Logs, string(p))
	// 处理日志数量
	if len(mlog.Logs) > mlog.MaxNum {
		mlog.Logs = mlog.Logs[len(mlog.Logs)-mlog.MaxNum:]
	}
	return len(p), nil
}

var mlogs = &memLogs{MaxNum: 50}

// 初始化日志
func init() {
	log.SetOutput(io.MultiWriter(mlogs, os.Stdout))
}

// logs web
func Logs(writer http.ResponseWriter, request *http.Request) {
	for _, log := range mlogs.Logs {
		writer.Write([]byte(log))
		writer.Write([]byte("<br/>"))
	}

}

// clear logs
func ClearLog(writer http.ResponseWriter, request *http.Request) {
	mlogs.Logs = []string{}
	writer.Write([]byte("日志已清空"))
}
