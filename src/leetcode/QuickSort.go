package main

import "fmt"

// 快排序V1版本
func quickSort(nums []int) {
	quickSortInner(nums, 0, len(nums)-1)
}

func quickSortInner(nums []int, low int, high int) {
	if low >= high {
		return
	}
	// 每次都会使用排序一个数据
	j := partition(nums, low, high)
	quickSortInner(nums, low, j-1)
	quickSortInner(nums, j+1, high)
}

// v < w true
//private static boolean less(Comparable v, Comparable w) {
//if (v == w) return false;   // optimization when reference equals
//return v.compareTo(w) < 0;
//}

func partition(nums []int, low int, high int) int {
	// 双轴指针
	lowTemp := low + 1
	highTemp := high
	v := nums[low]
	for true {
		// find item on lo to swap
		for nums[lowTemp] < v {
			if lowTemp == high {
				break
			} else {
				lowTemp = lowTemp + 1
			}
		}
		// find item on hi to swap
		for nums[highTemp] > v {
			if highTemp == low { // redundant since a[lo] acts as sentinel
				break
			} else {
				highTemp = highTemp - 1
			}
		}
		// check if pointers cross
		if lowTemp >= highTemp {
			break
		}

		// 两个找到了正确位置
		nums[lowTemp], nums[highTemp] = nums[highTemp], nums[lowTemp]
		fmt.Println(nums)
	}
	// 其实这里已经相等了
	//now, a[lo .. j-1] <= a[j] <= a[j+1 .. hi]
	nums[highTemp], nums[low] = nums[low], nums[highTemp]
	fmt.Printf("low:%d,high:%d\n", lowTemp, highTemp)
	//  已经把数据放到了highTemp
	fmt.Println(nums)
	return highTemp
}

func main() {
	intItems := make([]int, 10)
	for i := 9; i >= 0; i-- {
		intItems[9-i] = i
	}
	quickSort(intItems)
	fmt.Println(intItems)
}
