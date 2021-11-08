package main

import "math"

func findMode(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	m := make(map[int]int, 0)
	var r []int
	findModeHelper(root, m)
	max := 0
	for _, i := range m {
		if i > max {
			max = i
		}
	}
	for index, i := range m {
		if i == max {
			r = append(r, index)
		}
	}
	return r
}

func findModeHelper(root *TreeNode, m map[int]int) {
	if nil == root {
		return
	}
	findModeHelper(root.Left, m)
	m[root.Val]++
	findModeHelper(root.Right, m)
}

func getMinimumDifference(root *TreeNode) int {
	var temp []int
	getMinimumDifferenceHelper(root, &temp)
	length := len(temp)
	min := math.MaxInt64
	// 使用后继节点
	for i := range temp {
		if i < length && (i+1) < length {
			min = Min(min, temp[i+1]-temp[i])
		}
	}
	return min
}
func getMinimumDifferenceHelper(root *TreeNode, temp *[]int) {
	if nil == root {
		return
	}
	getMinimumDifferenceHelper(root.Left, temp)
	*temp = append(*temp, root.Val)
	getMinimumDifferenceHelper(root.Right, temp)
}
