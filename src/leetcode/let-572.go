package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(a, b *TreeNode) bool {
	// 两个都为空
	if a == nil && b == nil {
		return true
	}
	// 一个为空另外一个不为空
	if a == nil || b == nil {
		return false
	}
	// 相等的话，则需要都相等
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}

//[1,1]
// [1] 问题error
//func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
//	if root == nil && subRoot == nil {
//		return true
//	} else if root != nil && subRoot != nil {
//		if root.Val != subRoot.Val {
//			left := isSubtree(root.Left, subRoot)
//			right := isSubtree(root.Right, subRoot)
//			return left || right
//		} else {
//			leftE := isSubtree(root.Left, subRoot.Left)
//			rightE := isSubtree(root.Right, subRoot.Right)
//			return leftE && rightE
//		}
//	}
//	return false
//}

// error
//[3,4,5,1,null,2]
//[3,1,2]
//func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
//	if root == nil && subRoot == nil {
//		return true
//	} else if root != nil && subRoot != nil {
//		if root.Val != subRoot.Val {
//			left := isSubtree(root.Left, subRoot)
//			right := isSubtree(root.Right, subRoot)
//			return left || right
//		} else if nil != root.Left && root.Left.Val == subRoot.Val {
//			return isSubtree(root.Left, subRoot)
//		} else if nil != root.Right && root.Right.Val == subRoot.Val {
//			return isSubtree(root.Right, subRoot)
//		} else {
//			leftE := isSubtree(root.Left, subRoot.Left)
//			rightE := isSubtree(root.Right, subRoot.Right)
//			return leftE && rightE
//		}
//	}
//	return false
//}
