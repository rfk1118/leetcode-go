package main

var tilt int

func findTilt(root *TreeNode) int {
	tilt = 0
	findTiltHelper(root)
	return tilt
}

func findTiltHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左边和
	left := findTiltHelper(root.Left)
	// 右边和
	right := findTiltHelper(root.Right)
	// 坡度
	tilt = tilt + abs(left, right)
	// 总和
	return left + right + root.Val
}
