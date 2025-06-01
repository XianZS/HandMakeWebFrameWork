package context

import "net/http"

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response: w,
		request:  r,
		Pattern:  r.Method,
		Method:   r.URL.Path,
	}
}
