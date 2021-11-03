package main

import "math"

//func isValidBST(root *TreeNode) bool {
//	return isValidBSTValidate(root)
//}

func isValidBSTValidate(root *TreeNode) bool {
	if root == nil {
		return true
	}
	validateLeft := isValidBSTValidate(root.Left)
	validateRight := isValidBSTValidate(root.Right)
	leftCompareR := false
	rightCompareR := false
	if root.Left != nil {
		leftCompareR = root.Val > let98findMax(root.Left)
	} else {
		leftCompareR = true
	}
	if root.Right != nil {
		rightCompareR = root.Val < let98findMin(root.Right)
	} else {
		rightCompareR = true
	}
	return validateLeft && validateRight && leftCompareR && rightCompareR
}

func let98findMin(root *TreeNode) int {
	if root.Left != nil {
		return let98findMin(root.Left)
	}
	return root.Val
}

func let98findMax(root *TreeNode) int {
	if root.Right != nil {
		return let98findMax(root.Right)
	}
	return root.Val
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

// 秀翻我了
func helper(root *TreeNode, lower, upper int) bool {
	// 如果为空的话返回为true
	if root == nil {
		return true
	}
	// 如果当前节点小于最大节点或者小于小的返回为false
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
