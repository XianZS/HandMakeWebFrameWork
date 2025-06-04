package trie

// applicationNode [注册结点] 返回*NodeTree，用于存储应用程序数据。
func (n *node) applicationNode(part string, data string) *node {
	/*
		个人见解：对于每个结点来说，它需要知道“自己在哪里？”和“自己是什么”
		也就是说，需要知道它当前的路径和它当前的数据。
		注册结点，那么就需要传入：
		part string : 当前结点的路径
		data string : 当前结点的数据
		返回*NodeTree，用于存储应用程序数据。
	*/
	return nil
}
