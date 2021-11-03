package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return l100helper(p, q)
}

func l100helper(p, q *TreeNode) bool {
	if nil != q && nil != p {
		leftHelper := l100helper(p.Left, q.Left)
		rightHelper := l100helper(p.Right, q.Right)
		return q.Val == p.Val && leftHelper && rightHelper
	} else if nil == q && nil == p {
		return true
	} else {
		return false
	}
}
