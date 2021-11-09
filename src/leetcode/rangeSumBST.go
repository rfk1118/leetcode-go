package main

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	if root == nil {
		return sum
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if nil == root {
			return
		}
		dfs(root.Left)
		if root.Val >= low && root.Val <= high {
			sum = sum + root.Val
		}
		dfs(root.Right)
	}
	dfs(root)
	return sum
}
