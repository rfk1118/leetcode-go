package main

type TreeNode94 struct {
	Val   int
	Left  *TreeNode94
	Right *TreeNode94
}

func inorderTraversal(root *TreeNode94) []int {
	var r []int
	inorder(root, &r)
	return r
}

func inorder(root *TreeNode94, result *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}
