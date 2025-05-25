package main

import (
	"fmt"
	"github.com/XianZS/HandMakeWebFrameWork/server"
	"testing"
)

// TestHTTP_Start : This is a test function for the HTTP server.
// 测试HTTP服务器的启动函数.
func TestHTTP_Start(t *testing.T) {
	h := server.NewHTTPServer(server.WithHTTPServerStop(nil))
	go func() {
		err := h.Start(":8080")
		if err != nil {
			panic("启动失败")
		}
	}()
	err := h.Stop()
	if err != nil {
		panic("关闭失败")
	}
	fmt.Println(err)
}
