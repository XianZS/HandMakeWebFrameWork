package context

import "net/http"

func (context *Context) GetRequest() *http.Request {
	return context.request
}
