package main

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var temp01 []*TreeNode
	var temp02 []*TreeNode
	var dfs func(r *TreeNode, t *[]*TreeNode)
	dfs = func(r *TreeNode, t *[]*TreeNode) {
		if nil == r {
			return
		}
		// 提升性能，因为当前节点已经是叶子了，不需要递归查看孩子了
		if r.Left == nil && r.Right == nil {
			*t = append(*t, r)
			return
		}
		dfs(r.Left, t)
		dfs(r.Right, t)
	}
	dfs(root1, &temp01)
	dfs(root2, &temp02)
	if len(temp01) != len(temp02) {
		return false
	}
	for i := range temp01 {
		if temp01[i].Val != temp02[i].Val {
			return false
		}
	}
	return true
}

func main() {
	node01 := &TreeNode{Val: 2}
	node02 := &TreeNode{Val: 1}
	leafSimilar(node01, node02)
}
