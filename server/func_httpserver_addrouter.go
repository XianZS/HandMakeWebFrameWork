/*
Package server

	# 注册路由
		addRouter(请求方法，请求路径，视图函数)
	# 注册路由的时机
		项目启动时统一注册路由，项目启动之后就不能再注册路由了
	# 注册的路由放在哪里？
		routers map[string]HandleFunc
*/
package server

import "fmt"

func (h *HTTPServer) addRouter(method string, pattern string, handleFunc HandleFunc) {
	// 构建唯一的key
	key := fmt.Sprintf("%s-%s", method, pattern)
	h.routers[key] = handleFunc
}
