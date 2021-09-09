package main

import "math"

type RatioNode struct {
	next *RatioNode
	v    int
}

func RemoveByRatio(head *RatioNode, a, b int) *RatioNode {
	if a < 1 || a > b {
		return head
	}
	n := 0
	current := head
	for current.next != nil {
		n++
		current = current.next
	}
	ceil := math.Ceil(float64(a * n / b))
	if ceil == 1 {
		return head.next
	}
	// 删除第n个数据
	if ceil > 1 {
		current = head
		ceil--
		for ceil != 1 {
			current = current.next
			ceil--
		}
		// 删除第2个元素的话就是头的next是第三个元素
		current.next = current.next.next
	}
	return head
}
