// Package server
// "视图函数类型"是用来处理请求的函数，它的参数是http.ResponseWriter和*http.Request
package server

import "github.com/XianZS/HandMakeWebFrameWork/context"

// HandleFunc 之前函数前面
// type HandleFunc func(w http.ResponseWriter, r *http.Request)
// 更改函数签名
type HandleFunc func(ctx *context.Context)
