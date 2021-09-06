package main

import "fmt"

func singleNumber(nums []int) int {
	er := 0
	for _, num := range nums {
		er = er ^ num
	}
	return er
}

// ^ 如果相对应位值相同，则结果为0，否则为1
// ~ 按位取反运算符翻转操作数的每一位，即0变成1，1变成0。
func findTwoSingleNumber(nums []int) {
	// 先找到a^b
	er := 0
	other := 0
	for _, num := range nums {
		er = er ^ num
	}
	// 找到第二个数
	rightOne := er & (^er + 1)
	for _, num := range nums {
		if (num & rightOne) != 0 {
			other = other ^ num
		}
	}
	fmt.Printf("other:%d,another:%d", other, er^other)
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 5}
	findTwoSingleNumber(nums)
}
