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
	*/
	// 1. 分割字符串 ["user","address","id"]
	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	fmt.Println("==========", parts, "==========")
	fmt.Println(parts)
	fmt.Println("==========", parts, "==========")
	// 2. 判断合法性
	for _, part := range parts {
		if part == "" {
			panic("路由不合法")
		}
	}
}
