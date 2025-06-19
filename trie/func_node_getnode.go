package trie

func (n *Node) getNode(part string) *Node {
	// node 的 children 属性不存在
	if n.children == nil {
		return nil
	}
	// node 的 children 属性存在
	child, ok := n.children[part]
	if !ok {
		return nil
	}
	return child
}
