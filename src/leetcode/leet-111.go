package main

func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
