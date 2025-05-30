package server

import (
	"fmt"
	"net/http"
)

func (h *HTTPServer) Start(addr string) error {
	/*
		## Means:
			启动服务器.
			Start the server.
		## How to start the server?
			默认监听8080端口.
			Listen the address that default port is 8080.s
	*/
	h.srv = &http.Server{
		Addr:    addr,
		Handler: h,
	}
	return h.srv.ListenAndServe()
}

func prRouters(h *HTTPServer) {
	/*
		## Means:
			打印路由信息.
			Print the routers.
		## How to print the routers?
			遍历routers map，打印路由信息.
			Iterate the routers map, print the routers.
	*/
	fmt.Println("===***=== Routers ===***===")
	for key, value := range h.routers {
		fmt.Printf("[%v] -> ", key)
		fmt.Println(value)
	}
	fmt.Println("===***=== ===***=== ===***===")
}
