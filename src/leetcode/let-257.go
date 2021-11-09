package main

import "strconv"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var r []string
	itoa := strconv.Itoa(root.Val)
	binaryTreePathsHelper(root, &itoa, &r)
	return r
}

func binaryTreePathsHelper(root *TreeNode, s *string, r *[]string) {
	if root.Left == nil && root.Right == nil {
		*r = append(*r, *s)
		return
	}
	if root.Left != nil {
		t := *s + "->" + strconv.Itoa(root.Left.Val)
		binaryTreePathsHelper(root.Left, &t, r)
	}
	if root.Right == nil {
		t := *s + "->" + strconv.Itoa(root.Right.Val)
		binaryTreePathsHelper(root.Right, &t, r)
	}
}

// [1,2,3,null,5]
func main() {
	one := &TreeNode{
		Val: 1,
	}
	two := &TreeNode{
		Val: 2,
	}
	three := &TreeNode{
		Val: 3,
	}
	five := &TreeNode{
		Val: 5,
	}
	one.Left = two
	one.Right = three
	two.Right = five
	binaryTreePaths(one)
}
