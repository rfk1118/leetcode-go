package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	//leftSearch := searchBST(root.Left, val)
	//rightSearch := searchBST(root.Right, val)
	//if leftSearch == nil {
	//	return rightSearch
	//}
	//return leftSearch
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
