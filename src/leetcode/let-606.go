package main

import (
	"fmt"
	"strconv"
)

// 如果当前节点有两个孩子，那我们在递归时，需要在两个孩子的结果外都加上一层括号；
func tree2str(root *TreeNode) string {
	if nil == root {
		return ""
	}
	// 如果当前节点没有孩子，那我们不需要在节点后面加上任何括号；
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val) + ""
	}
	// 如果当前节点只有左孩子，那我们在递归时，只需要在左孩子的结果外加上一层括号，而不需要给右孩子加上任何括号；
	if root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
	}
	// 如果当前节点只有右孩子，那我们在递归时，需要先加上一层空的括号 () 表示左孩子为空，
	// 再对右孩子进行递归，并在结果外加上一层括号。
	return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")" + "(" + tree2str(root.Right) + ")"
}

func main() {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node4 := &TreeNode{
		Val: 4,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4

	str := tree2str(node1)
	fmt.Println(str)
}
