/*
Package server
向前对接前端请求，向后对接框架
|-----前端-------|       |---------后端--------|
|               |       |     框架<-------|   |
|               |       |                |   |
|               |---|   |————————————|   |   |
|               |   |   |            |   |   |
|               |   |-->| HTTPServer |---|   |
|               |       |            |       |
|               |       |            |       |
|               |       |------------|       |
|               |       |                    |
|---------------|       |--------------------|
*/

package server

import (
	"fmt"
	"github.com/XianZS/HandMakeWebFrameWork/context"
	"net/http"
)

func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*
		# Means:
			接收和转发请求函数.
			Receive the request and forward the request.
		# Receive the request:
			接收一个请求.
			* Receive the request from the client.
		# Forward the request:
			转发一个请求.
			* Forward the request to the router.
	*/
	// 1. 匹配路由
	key := fmt.Sprintf("%s-%s", r.Method, r.URL.Path)
	handler, ok := h.routers[key] // 返回视图函数
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404 NOT FOUND"))
		return
	}
	// 2. 构造当前请求的上下文
	c := context.NewContext(w, r)
	// 2. 转发请求
	handler(c)
}
