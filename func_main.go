package main

import (
	"errors"
	"fmt"
	"github.com/XianZS/HandMakeWebFrameWork/server"
	"net/http"
)

func main() {
	h := server.NewHTTPServer(server.WithHTTPServerStop(nil))
	go func() {
		err := h.Start(":8080")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("启动失败")
		}
	}()
	err := h.Stop()
	if err != nil {
		panic("关闭失败")
	}
	fmt.Println(err)
}
