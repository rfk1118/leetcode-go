package main

func bubbleSort(num []int) []int {
	length := len(num)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length-i; j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	return num
}
