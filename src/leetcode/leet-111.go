package main

import "math"

func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return minHDepth(root)
}

func minHDepth(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	midD := math.MaxInt64
	// 又一个不为空就需要向下走
	if root.Left != nil {
		midD = min(midD, minDepth(root.Left))
	}
	if root.Right != nil {
		midD = min(minDepth(root.Right), midD)
	}
	return midD + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
