package context

import "net/http"

// Context : 上下文
type Context struct {
	// w 响应
	response http.ResponseWriter
	// r 请求
	request *http.Request
	// Pattern 请求路径
	Pattern string
	// Method 请求方法
	Method string
}
