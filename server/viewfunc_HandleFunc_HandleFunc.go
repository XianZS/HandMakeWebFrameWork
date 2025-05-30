// Package server
// "视图函数类型"是用来处理请求的函数，它的参数是http.ResponseWriter和*http.Request
package server

import "net/http"

type HandleFunc func(w http.ResponseWriter, r *http.Request)
