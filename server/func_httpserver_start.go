package server

import (
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
