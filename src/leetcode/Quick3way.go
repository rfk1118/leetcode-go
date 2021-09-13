package main

import "fmt"

// QuickSort3Way 切割成 x<T.....T<y
func QuickSort3Way(num []int) {
	QuickSortWithIndex3Way(num, 0, len(num)-1)
}

// QuickSortWithIndex3Way 对从low~high处的数据进行排序
func QuickSortWithIndex3Way(num []int, low, high int) {
	if low >= high {
		return
	}
	// 设置双轴
	lt := low
	gt := high
	pV := num[low]
	index := low + 1
	for index <= gt {
		// index向下走
		if num[index] < pV {
			num[lt], num[index] = num[index], num[lt]
			lt++
			index++
		} else if num[index] > pV {
			num[gt], num[index] = num[index], num[gt]
			gt--
		} else {
			index++
		}
	}
	//partition := partitionWithIndex(num, low, high)
	QuickSortWithIndex3Way(num, low, lt-1)
	QuickSortWithIndex3Way(num, gt+1, high)
}

func main() {
	ns := make([]int, 0)
	for i := 1; i < 10; i++ {
		ns = append(ns, i)
	}
	QuickSort3Way(ns)
	fmt.Println(ns)
}
