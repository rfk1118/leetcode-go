package main

var temp []int

func findTarget(root *TreeNode, k int) bool {
	temp = make([]int, 0)
	findTargetHelper(root)
	minIndex, maxIndex := 0, len(temp)-1
	for minIndex < maxIndex {
		if temp[minIndex]+temp[maxIndex] == k {
			return true
		} else if temp[minIndex]+temp[maxIndex] < k {
			minIndex++
		} else {
			maxIndex--
		}
	}
	return false
}

func findTargetHelper(root *TreeNode) {
	if nil == root {
		return
	}
	findTargetHelper(root.Left)
	temp = append(temp, root.Val)
	findTargetHelper(root.Right)
}
