package trie

import (
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
			wantRouter: &Router{map[string]*node{
				"/": {
					part: "/",
					children: map[string]*node{
						"user": {
							part: "user",
							children: map[string]*node{
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
	router := &Router{map[string]*node{
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
