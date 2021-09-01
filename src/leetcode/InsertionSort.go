package main

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
