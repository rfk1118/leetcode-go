package main

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return abs(height(root.Right), height(root.Left)) <= 1 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return Max(height(root.Left), height(root.Right)) + 1
}

func abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
