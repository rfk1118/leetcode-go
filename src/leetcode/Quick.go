package main

import "fmt"

// QuickSort 标准快排，这里使用lo元素作为切割垫
func QuickSort(num []int) {
	QuickSortWithIndex(num, 0, len(num)-1)
}

// QuickSortWithIndex 对从low~high处的数据进行排序
func QuickSortWithIndex(num []int, low, high int) {
	if low >= high {
		return
	}
	partition := partitionWithIndex(num, low, high)
	QuickSortWithIndex(num, low, partition-1)
	QuickSortWithIndex(num, partition+1, high)
}

//  核心代码在这里也就是进行双轴 partition
func partitionWithIndex(num []int, low int, high int) int {
	lowIndex := low + 1
	highIndex := high
	pV := num[low]
	// 双轴交换大小值
	for true {
		for num[lowIndex] < pV {
			if lowIndex >= high {
				break
			} else {
				lowIndex++
			}
		}

		for num[highIndex] > pV {
			if highIndex <= low {
				break
			} else {
				highIndex--
			}
		}
		if lowIndex >= highIndex {
			break
		}
		num[lowIndex], num[highIndex] = num[highIndex], num[lowIndex]
	}
	num[low], num[highIndex] = num[highIndex], num[low]
	return highIndex
}

func main() {
	ns := make([]int, 0)
	for i := 1; i < 10; i++ {
		ns = append(ns, i)
	}
	QuickSort(ns)
	fmt.Println(ns)
}
