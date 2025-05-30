package server

func NewHTTPServer(opts ...HttpOption) *HTTPServer {
	// 对于go语言来说，每一个结构体都有一个go方法，这个方法就是NewXXX
	h := &HTTPServer{
		routers: map[string]HandleFunc{},
	}
	// option 模式 : 可以在创建对象的时候，传入一些参数，来改变对象的行为
	for _, opt := range opts {
		opt(h)
	}
	return h
}
