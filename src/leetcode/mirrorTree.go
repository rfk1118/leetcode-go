package main

func mirrorTree(root *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	var bfs func(node *TreeNode)
	bfs = func(node *TreeNode) {
		if nil == node {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		bfs(node.Left)
		bfs(node.Right)
	}
	bfs(root)
	return root
}
