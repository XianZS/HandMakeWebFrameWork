package trie

import (
	"errors"
	"strings"
)

/*
GetRouter [获取路由]
公开性质:public
*/
func (r *Router) GetRouter(pattern string) (*Node, error) {
	// 找到pattern对应的node，也就是对应的根节点
	// 先寻找根路由
	root, ok := r.root["/"]
	// 创建跟路由
	if !ok {
		// 根节点不存在
		return nil, errors.New("root not found")
	}
	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	for _, part := range parts {
		// 判断root不能为nil，并且part值合法
		if part == "" {
			return nil, errors.New("pattern格式不正确")
		}
		root = root.getNode(part)
		if root == nil {
			// fmt.Println("-=-=-=\('*-*')/-=--=-=")
			return nil, errors.New("pattern不存在")
		}
	}
	// fmt.Println("-=-=-=\('*-*')/-=--=-=")
	return root, nil
}
