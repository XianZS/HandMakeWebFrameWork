package main

import (
	"fmt"
	"github.com/XianZS/HandMakeWebFrameWork/server"
	"net/http"
	"testing"
)

// TestHTTP_Start_1 : This is a test function for the HTTP server.
// 测试HTTP服务器的启动函数.
// 2025-5-25
func TestHTTP_Start_1(t *testing.T) {
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

// TestHTTP_Start_2
// 测试GET、POST、PUT、DELETE方法的路由注册.
// 2025-5-30
func TestHTTP_Start_2(t *testing.T) {
	h := server.NewHTTPServer()
	// 注册路由 : 在启动之前
	// 接口测试
	h.GET("/world", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello world"))
	})
	h.GET("/user", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello user"))
	})
	h.POST("/user", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("POST:hello user"))
	})
	// 启动服务器
	err := h.Start(":8080")
	if err != nil {
		panic("启动失败")
	}
}
