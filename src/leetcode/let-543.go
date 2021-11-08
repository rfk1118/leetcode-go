package main

var diameter int

// 本题核心在于diameter = Max(diameter, Lmax+Rmax)
// 这里会记录最大值，因为经过root的不一定获取到最大的结果
func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	diameterOfBinaryTreeHelper(root)
	return diameter
}

func diameterOfBinaryTreeHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 只计算左边的深度
	Lmax := diameterOfBinaryTree(root.Left)
	// 只计算右边的深度
	Rmax := diameterOfBinaryTree(root.Right)
	// 记录最大，
	diameter = Max(diameter, Lmax+Rmax)
	// 记录深度
	return Max(Lmax, Rmax) + 1
}
