package server

import "net/http"

/* ========= 结构体对象定义 ========= */

// HTTPServer : This is a http server.
type HTTPServer struct {
	srv  *http.Server
	stop func() error
}
