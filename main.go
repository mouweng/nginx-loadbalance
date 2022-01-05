package main

import (
	"net/http"
	"os"
	"strings"
)

var post string = ":3003"

func main() {
	http.HandleFunc("/test", handleReq)
	http.ListenAndServe(post, nil)
}

//处理请求函数,根据请求将响应结果信息写入日志
func handleReq(w http.ResponseWriter, r *http.Request) {
	failedMsg := "handle in port:"
	writeLog(failedMsg, "./stat.log")
}

//写入日志
func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, post)
	buf := []byte(content)
	fd.Write(buf)
}
