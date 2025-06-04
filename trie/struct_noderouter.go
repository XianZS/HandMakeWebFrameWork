package trie

/*
NodeRouter [根节点] 提供给用户。
公开性质:public
*/

type NodeRouter struct {
	root map[string]*NodeRouter
}
