package main

func convertBiNode(root *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	t := &TreeNode{}
	c := t
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}
		dfs(node.Left)
		c.Right = node
		node.Left = nil
		c = node
		dfs(node.Right)
	}
	dfs(root)
	return t.Right
}
