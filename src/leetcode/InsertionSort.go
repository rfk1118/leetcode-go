package main

import "fmt"

// 直接插入排序（Straight Insertion Sort）是一种最简单的排序方法，其基本操作是将一条记录插入到已排好序的有序表中，从而得到一个新的、记录数量增1的有序表。
func insertSort(num []int) []int {
	// 总长度
	length := len(num)
	// 不能从0开始因为第一个元素要跟0对比
	for i := 1; i < length; i++ {
		for j := i; j > 0 && num[j] < num[j-1]; j-- {
			num[j] = num[j] ^ num[j-1]
			num[j-1] = num[j] ^ num[j-1]
			num[j] = num[j] ^ num[j-1]
		}
	}
	return num
}

//
func insertSortWithFront(num []int) []int {
	length := len(num)
	// 不能从0开始因为第一个元素要跟0对比
	for i := 0; i < length; i++ {
		successIndex := i
		insertV := num[i]
		for j := 0; j < i; j++ {
			if num[j] > num[i] {
				successIndex = j
				break
			}
		}
		for j := i; j > successIndex; j-- {
			num[j] = num[j-1]
		}
		num[successIndex] = insertV
	}
	return num
}

func main() {
	intItems := make([]int, 10)
	for i := 9; i >= 0; i-- {
		intItems[9-i] = i
	}
	insertSortWithFront(intItems)
	fmt.Println(intItems)
}
