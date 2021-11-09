package main

import "math"

func minDiffInBST(root *TreeNode) int {
	r, pre := math.MaxInt64, -1
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < r {
			r = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return r
}
