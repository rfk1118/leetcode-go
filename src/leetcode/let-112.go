package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumHelper(root, 0, targetSum)
}

func hasPathSumHelper(root *TreeNode, current, targetSum int) bool {
	if nil == root {
		return false
	}
	if nil == root.Left && nil == root.Right {
		return current+root.Val == targetSum
	}
	current = root.Val + current
	return hasPathSumHelper(root.Left, current, targetSum) ||
		hasPathSumHelper(root.Right, current, targetSum)
}
