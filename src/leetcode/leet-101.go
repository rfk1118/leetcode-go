package main

func isSymmetric(root *TreeNode) bool {
	return l101helper(root.Left, root.Right)
}

func l101helper(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	} else if q != nil && p != nil {
		if p.Val != q.Val {
			return false
		}
		return l101helper(p.Right, q.Left) && l101helper(p.Left, q.Right)
	} else {
		return false
	}
}
