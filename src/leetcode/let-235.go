package main

// 自底而上
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, q, p)
	right := lowestCommonAncestor(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// 利用二叉搜索树的概念
func lowestCommonAncestor02(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

func lowestCommonAncestor03(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	var pPath []*TreeNode
	var qPath []*TreeNode
	lowestCommonAncestorHelper(root, p, &pPath)
	lowestCommonAncestorHelper(root, q, &qPath)
	result := root
	for i := 0; i < len(pPath) && i < len(qPath) && pPath[i].Val == qPath[i].Val; i++ {
		result = pPath[i]
	}
	return result
}

func lowestCommonAncestorHelper(root, p *TreeNode, path *[]*TreeNode) {
	*path = append(*path, root)
	if p.Val == root.Val {
		return
	}
	if p.Val > root.Val {
		lowestCommonAncestorHelper(root.Right, p, path)
	}
	if p.Val < root.Val {
		lowestCommonAncestorHelper(root.Left, p, path)
	}
}
