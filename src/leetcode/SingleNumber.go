package main

func singleNumber(nums []int) int {
	error := 0
	for _, num := range nums {
		error = error ^ num
	}
	return error
}
