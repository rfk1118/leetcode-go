package main

import "fmt"

// 处理小和问题
func minMergeSort(num []int) int {
	aux := make([]int, len(num))
	return mergeSortWithAux(num, aux, 0, len(num)-1)
}

func mergeSortWithAux(num []int, aux []int, low int, high int) int {
	if low >= high {
		return 0
	}
	mid := low + (high-low)/2
	lowCount := mergeSortWithAux(num, aux, low, mid)
	highCount := mergeSortWithAux(num, aux, mid+1, high)
	return lowCount + highCount + mergeWithAux(num, aux, low, mid, high)
}

func mergeWithAux(num []int, aux []int, low int, mid int, high int) int {
	lowIndex := low
	highIndex := mid + 1

	for i := low; i <= high; i++ {
		aux[i] = num[i]
	}

	allCount := 0
	for i := low; i <= high; i++ {
		if lowIndex > mid {
			num[i] = aux[highIndex]
			highIndex = highIndex + 1
		} else if highIndex > high {
			num[i] = aux[lowIndex]
			lowIndex = lowIndex + 1
		} else if aux[lowIndex] < aux[highIndex] {
			allCount += aux[lowIndex] * (high - highIndex + 1)
			num[i] = aux[lowIndex]
			lowIndex = lowIndex + 1
		} else {
			num[i] = aux[highIndex]
			highIndex = highIndex + 1
		}
	}
	return allCount
}

func main() {
	num := []int{1, 3, 5, 2, 4, 6}
	sort := minMergeSort(num)
	fmt.Println(sort)
}
