package main

func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var res = make([][]int, 0)
	q := []*TreeNode{root}
	for len(q) != 0 {
		length := len(q)
		var t []int
		for i := 0; i < length; i++ {
			node := q[i]
			t = append(t, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, t)
		q = q[length:]
	}
	return res
}
