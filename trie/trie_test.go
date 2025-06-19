package trie

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouter_AddRouter(t *testing.T) {
	testCases := []struct {
		name       string
		pattern    string
		data       string
		wantRouter *Router
	}{
		{
			name:    "xxx",
			pattern: "/user/login",
			data:    "hello",
			wantRouter: &Router{map[string]*Node{
				"/": {
					part: "/",
					children: map[string]*Node{
						"user": {
							part: "user",
							children: map[string]*Node{
								"login": {
									part: "login",
									data: "hello",
								},
							},
							data: "hello",
						},
					},
					data: "ddd",
				},
			}},
		},
	}
	router := &Router{map[string]*Node{
		"/": {
			part: "/",
		},
	}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router.AddRouter(tc.pattern, tc.data)
			assert.Equal(t, tc.wantRouter.root["/"].children["user"].children["login"].data, "hello")
		})
	}
}

func TestRouter_GetRouter(t *testing.T) {
	testCases := []struct {
		// 测试的名字 任意给
		name string
		// 想要匹配的路由结点
		findPattern string
		// 想要返回的数据
		wantData string
		// 理想中的错误
		wantErr error
	}{
		// test_1: 正确测试
		{
			name:        "success",
			findPattern: "/user/login",
			wantData:    "hello",
		},
		// test_2: pattern 格式不正确
		{
			name:        "pattern 格式不正确",
			findPattern: "/user//login",
			wantErr:     errors.New("pattern格式不正确"),
		},
		// test_3: node 不存在
		{
			name:        "node 不存在",
			findPattern: "/user/login/m",
			wantErr:     errors.New("pattern不存在"),
		},
	}
	router := &Router{map[string]*Node{
		"/": {
			part: "/",
		},
	}}
	router.AddRouter("/user/login", "hello")
	router.AddRouter("/user/register", "world")
	router.AddRouter("/study/golang", "Good")
	router.AddRouter("/study/python", "Py")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n, err := router.GetRouter(tc.findPattern)
			fmt.Println("-------------==================------------------")
			fmt.Println(tc.wantErr, err)
			fmt.Println("-------------==================------------------")
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantData, n.data)
		})
	}
}
