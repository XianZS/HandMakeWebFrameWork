package trie

import (
	"errors"
	"strings"
)

/*
GetRouter [获取路由]
公开性质:public
*/
func (r *Router) GetRouter(pattern string) (*node, error) {
	// 获取根路由
	root, ok := r.root["/"]
	// 创建跟路由
	if !ok {
		// 根节点不存在
		return nil, errors.New("root not found")
	}
	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	for _, part := range parts {
		// 获取根路由
		if part == "" {
			return nil, errors.New("pattern 格式不正确")
		}
		root = root.getnode(part)
	}
	return
}
