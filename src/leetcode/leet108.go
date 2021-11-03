package main

func sortedArrayToBST(nums []int) *TreeNode {
	if nil == nums || len(nums) == 0 {
		return nil
	}
	return addNode(nums, 0, len(nums)-1)
}

func addNode(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := (high-low)/2 + low
	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	root.Left = addNode(nums, low, mid-1)
	root.Right = addNode(nums, mid+1, high)
	return root
}
