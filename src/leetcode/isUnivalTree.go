package main

func isUnivalTree(root *TreeNode) bool {
	if nil == root {
		return true
	}
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if nil == node {
			return true
		}
		leftB := dfs(node.Left)
		rightB := dfs(node.Right)
		return leftB && rightB && node.Val == root.Val
	}
	return dfs(root)
}
