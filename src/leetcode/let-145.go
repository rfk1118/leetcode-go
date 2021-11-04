package main

func postorderTraversal(root *TreeNode) []int {
	var r []int
	postorderTraversalHelper(root, &r)
	return r
}

func postorderTraversalHelper(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}
	postorderTraversalHelper(root.Left, r)
	postorderTraversalHelper(root.Right, r)
	*r = append(*r, root.Val)
}
