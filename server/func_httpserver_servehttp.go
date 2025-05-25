package server

import "net/http"

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
	// TODO: Implement me!
	panic("implement me")
}
