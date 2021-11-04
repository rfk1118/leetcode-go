package main

func preorderTraversal(root *TreeNode) []int {
	var r []int
	preorderTraversalHelper(root, &r)
	return r
}

func preorderTraversalHelper(root *TreeNode, r *[]int) {
	if root == nil {
		return
	} else {
		*r = append(*r, root.Val)
	}
	preorderTraversalHelper(root.Left, r)
	preorderTraversalHelper(root.Right, r)
}
