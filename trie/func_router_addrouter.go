package trie

import (
	"fmt"
	"strings"
)

/*
AddRouter [添加路由]
公开性质:public
*/
func (r *Router) AddRouter(pattern string, data string) {
	/*
		个人见解：
		str_s :=/user/address/id
		对str_s进行分割，得到一个字符串数组。
		将分割之后的结果，分块保存到前缀树上。
		先默认用户传的数据都是标准的
		逻辑处理：
			先判断有没有根路由，如果有就使用，没有就去注册
	*/
	// 0. 预[处理/判断]阶段
	if r.root == nil {
		// 人话：处理Trie树为空的情况
		// 处理r.root未初始化的情况
		r.root = make(map[string]*node)
	}
	root, ok := r.root["/"]
	if !ok {
		// 创建根root结点
		root = &node{
			part: "/",
		}
		r.root["/"] = root
	}
	// 1. 分割字符串 ["user","address","id"]
	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	fmt.Println("==========", parts, "==========")
	fmt.Println(parts)
	fmt.Println("==========", parts, "==========")
	// 2. 判断合法性
	for _, part := range parts {
		if part == "" {
			// eg: /user//address
			// parts : ["user","","address"]
			panic("路由不合法")
		}
	}
}
