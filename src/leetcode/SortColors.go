package main

import "fmt"

func sortColors(nums []int) {
	sortCol(nums, 0, len(nums)-1)
}

func sortCol(nums []int, low, high int) {
	if low >= high {
		return
	}
	lowIndex := low
	index := low + 1
	highIndex := high
	compareV := nums[lowIndex]
	for index <= highIndex {
		if nums[index] < compareV {
			nums[lowIndex], nums[index] = nums[index], nums[lowIndex]
			lowIndex++
			index++
		} else if nums[index] > compareV {
			nums[index], nums[highIndex] = nums[highIndex], nums[index]
			highIndex--
		} else {
			index++
		}
	}
	sortCol(nums, low, lowIndex-1)
	sortCol(nums, highIndex+1, high)
}

func main() {
	num := []int{2, 0, 2, 1, 1, 0}
	sortColors(num)
	fmt.Println(num)
}
