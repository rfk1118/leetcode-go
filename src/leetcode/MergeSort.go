package main

func sortArray(num []int) []int {
	mergeSort(num, 0, len(num)-1)
	return num
}

// 从小到大
func mergeSort(nums []int, low, high int) {
	if low == high {
		return
	}
	mid := low + (high-low)/2
	mergeSort(nums, low, mid)
	mergeSort(nums, mid+1, high)
	merge(nums, mid, low, high)
}

func merge(nums []int, mid, low, high int) {
	leftMin := low
	rightMin := mid + 1

	num := make([]int, high-low+1)
	numIndex := 0

	for leftMin <= mid && rightMin <= high {
		if nums[leftMin] <= nums[rightMin] {
			num[numIndex] = nums[leftMin]
			leftMin = leftMin + 1
		} else {
			num[numIndex] = nums[rightMin]
			rightMin = rightMin + 1
		}
		numIndex = numIndex + 1
	}
	for leftMin <= mid {
		num[numIndex] = nums[leftMin]
		leftMin = leftMin + 1
		numIndex = numIndex + 1
	}
	for rightMin <= high {
		num[numIndex] = nums[rightMin]
		rightMin = rightMin + 1
		numIndex = numIndex + 1
	}
	for i, v := range num {
		nums[low+i] = v
	}
}
