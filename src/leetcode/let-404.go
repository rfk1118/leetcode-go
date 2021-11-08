package main

// 从顶向下找左子树
func sumOfLeftLeaves(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return bfs(root)
}

// 是否为叶子节点
func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func bfs(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	// 切片
	q := []*TreeNode{root}
	// 非空
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			if isLeaf(node.Left) {
				ans = ans + node.Left.Val
			} else {
				q = append(q, node.Left)
			}
			if node.Right != nil && !isLeaf(node.Right) {
				q = append(q, node.Right)
			}
		}
	}
	return
}

func dfs(root *TreeNode) (ans int) {
	if root.Left != nil {
		// 找到了左叶子节点
		if isLeaf(root.Left) {
			ans += root.Left.Val
		} else {
			// 还没找到，继续递归去找
			ans += dfs(root.Left)
		}
	}
	// 处理右节点
	if root.Right != nil && !isLeaf(root.Right) {
		ans += dfs(root.Right)
	}
	return
}
