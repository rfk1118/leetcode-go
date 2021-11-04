package main

func invertTree(root *TreeNode) *TreeNode {
	invertTreeHelper(root)
	return root
}
func invertTreeHelper(root *TreeNode) {
	if nil == root {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invertTreeHelper(root.Left)
	invertTreeHelper(root.Right)
}
