package context

import "net/http"

// GetResponse : 获取响应
func (context *Context) GetResponse() http.ResponseWriter {
	return context.response
}
