package server

import "net/http"

/* ========= 结构体对象定义 ========= */

// HTTPServer : This is a http server.
type HTTPServer struct {
	srv *http.Server
	// 优雅关闭服务
	stop func() error
	/* routers 临时存放路由信息
	{
		"GET /": func(w http.ResponseWriter, r *http.Request) { // 处理请求 },
		"POST /": func(w http.ResponseWriter, r *http.Request) { // 处理请求 },
	}
	*/
	routers map[string]HandleFunc
}
