package main

// 选择排序
func selectSort(num []int) []int {
	length := len(num)
	// 从第0开始查找最小的值
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if num[j] < num[minIndex] {
				minIndex = j
			}
		}
		// 交换
		temp := num[i]
		num[i] = num[minIndex]
		num[minIndex] = temp
	}
	return num
}
