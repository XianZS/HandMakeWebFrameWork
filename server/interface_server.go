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
	// Start : This is start server.
	Start(addr string) error
	// Stop : This is stop server
	Stop() error
}
