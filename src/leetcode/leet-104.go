package main

func maxDepth(root *TreeNode) int {
	return l104helper(root)
}

func l104helper(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	left := l104helper(root.Left)
	right := l104helper(root.Right)
	if left > right {
		result = left
	} else {
		result = right
	}
	return result + 1
}
