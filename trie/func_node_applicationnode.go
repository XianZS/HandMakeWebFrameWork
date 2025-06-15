package trie

// applicationNode [注册结点] 返回*NodeTree，用于存储应用程序数据。
// func (n *node) applicationNode(part string, data string) *node {
func (n *node) applicationNode(part string) *node {
	/*
		[注册结点]
		个人见解：对于每个结点来说，它需要知道“自己在哪里？”和“自己是什么”
		对于每条url来说，我们通过叶子结点来给当前路径打上唯一标识。
		比如说，对于/user/a1和/user/a2来说，它们有相同父结点，
		想要标识两条路径，只需要给叶子结点赋值即可，node->a1，node->a2。
		至于返回值，(思考‘……’)，当添加了一个结点之后，
		我们有可能需要添加下一个它的子结点，所以需要返回当前结点。
		---------
		也就是说，需要知道它当前的路径和它当前的数据。
		注册结点，那么就需要传入：
		part string : 当前结点的路径
		data string : 当前结点的数据
		返回*NodeTree，用于存储应用程序数据。
	*/
	// [没有] - 判断当前结点有没有
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	// [有] - 判断当前结点有没有
	child, ok := n.children[part]
	if !ok {
		child = &node{
			part: part,
		}
		n.children[part] = child
	}
	return child
}
