package main

var lRoot *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var temp []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		temp = append(temp, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	dummyNode := &TreeNode{}
	current := dummyNode
	for i := range temp {
		node := &TreeNode{Val: temp[i]}
		current.Right = node
		current = node
	}
	return dummyNode.Right
}

func increasingBST02(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	pre := dummyNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		// 前驱节点的右节点为当前节点
		pre.Right = root
		// 左节点设置为空
		root.Left = nil
		// 更新前驱节点
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	return dummyNode.Right
}
