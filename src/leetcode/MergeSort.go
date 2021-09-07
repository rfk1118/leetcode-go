package main

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

func sortArray(num []int) []int {
	aux := make([]int, len(num))
	mergeTopDownSort(num, aux, 0, len(num)-1)
	//mergeSort(num, 0, len(num)-1)
	return num
}

// aux 辅助数组
func mergeTopDownSort(nums []int, aux []int, low, high int) {
	if low >= high {
		return
	}
	mid := low + (high-low)/2
	mergeTopDownSort(nums, aux, low, mid)
	mergeTopDownSort(nums, aux, mid+1, high)
	mergeTopDown(nums, aux, low, mid, high)
}

func mergeTopDown(nums []int, aux []int, low, mid, high int) {
	// copy data to nums
	for i := low; i <= high; i++ {
		aux[i] = nums[i]
	}

	lowIndex := low
	highIndex := mid + 1
	for i := low; i <= high; i++ {
		if lowIndex > mid {
			//  左边获取完毕
			nums[i] = aux[highIndex]
			highIndex = highIndex + 1
		} else if highIndex > high {
			// 右边获取完毕
			nums[i] = aux[lowIndex]
			lowIndex = lowIndex + 1
		} else if aux[lowIndex] < aux[highIndex] {
			// 都没获取完毕，并且左边小于右边
			nums[i] = aux[lowIndex]
			lowIndex = lowIndex + 1
		} else {
			// 都没获取完毕，并且右边小于左边
			nums[i] = aux[highIndex]
			highIndex = highIndex + 1
		}
	}
}

// 自底而上处理
func mergeDownTopSort(num []int) {
	length := len(num)
	aux := make([]int, length)
	for sz := 1; sz < length; sz = sz + sz { //sz 子数组大小
		//  因为每次合并的都是2sz,所以下次处理的位置是2sz
		// 1 2 4 8 16
		for low := 0; low < length-sz; low += sz + sz { //子数组索引
			high := low
			if low+sz+sz-1 > length-1 {
				high = length - 1
			} else {
				high = low + sz + sz - 1
			}
			mergeTopDown(num, aux, low, low+sz-1, high)
		}
	}
}
