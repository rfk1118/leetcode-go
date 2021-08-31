package main

// 升序 最主要的是处理好角标
// 大兄弟长期不写算法，二分都费劲了哈哈哈
func searchWithLoop(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func searchWithRecursion(nums []int, target int) int {
	return searchWithRec(nums, target, 0, len(nums)-1)
}

func searchWithRec(nums []int, target int, low int, high int) int {
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			return searchWithRec(nums, target, mid+1, high)
		} else {
			return searchWithRec(nums, target, low, mid-1)
		}
	}
	return -1
}
