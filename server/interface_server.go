package server

import (
	"net/http"
)

/* ========= 接口定义 ========= */
type server interface {
	/*
		## Means
			make goland's http package intake server's interface_defined
		## Server's function?
			* Start
			* Stop
		## Why we need to abstract the http server?
			* We can use the http server to implement the server interface_defined
			* We need compatible with the http server
	*/
	http.Handler
	// Start : 开启服务
	Start(addr string) error
	// Stop : 关闭服务
	Stop() error
	// addRouter : 注册路由 addRouter(请求方法，请求路径，视图函数)
	/*
		视图函数是用来处理请求的函数，它的参数是http.ResponseWriter和*http.Request
	*/
	addRouter(method string, pattern string, handleFunc HandleFunc)
}
