package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// HttpOption : option 模式
type HttpOption func(h *HTTPServer)

// WithHTTPServerStop : option 模式
func WithHTTPServerStop(fn func() error) HttpOption {
	return func(h *HTTPServer) {
		if fn == nil {
			fn = func() error {
				fmt.Println("-0-0-0-0-0-0-0-0-0-0-0-0-")
				quit := make(chan os.Signal)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit
				fmt.Println("Shutdown Server ...")
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				// 关闭之前的操作
				if err := h.srv.Shutdown(ctx); err != nil {
					fmt.Println("Server Shutdown:", err)
					return err
				}
				// 关闭之后的操作
				select {
				case <-ctx.Done():
					log.Println("timeout of 5seconds.")
				}
				return nil
			}
		}
		h.stop = fn
	}
}
