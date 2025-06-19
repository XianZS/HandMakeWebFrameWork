package trie

/*
Node [结点] 用于存储应用程序数据。
公开性质:private
*/
type Node struct {
	// part [维护当前结点] 当前结点的唯一标识
	part string
	// children [维护子节点] 当前结点的子节点 map || slice
	children map[string]*Node
	// data [结点数据] 存储当前结点的数据
	data string
}
