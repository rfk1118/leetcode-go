package main

// 从顶向下找左子树
func sumOfLeftLeaves(root *TreeNode) int {
	var all []int
	sumOfLeftLeavesHelper(root, &all)
	sum := 0
	for _, i := range all {
		sum = sum + i
	}
	return sum
}

func sumOfLeftLeavesHelper(root *TreeNode, all *[]int) {
	// 往右走，已经为叶子节点，不在往下走了
	if root.Left == nil && root.Right == nil {
		return
	}
	node := sumOfFindMin(root)
	if nil != node {
		*all = append(*all, node.Val)
	}
	if nil != root.Right {
		sumOfLeftLeavesHelper(root.Right, all)
	}
}

// 找到所有的叶子节点
func sumOfFindMin(root *TreeNode) *TreeNode {
	// 没有叶子节点
	if root == nil {
		return nil
	}
	// 这里说明是叶子节点
	if root.Left == nil && root.Right == nil {
		return root
	}
	return sumOfFindMin(root.Left)
}
