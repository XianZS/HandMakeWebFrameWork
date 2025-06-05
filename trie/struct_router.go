package trie

/*
Router [根节点] 提供给用户。
公开性质:public
*/

type Router struct {
	root map[string]*node
}
