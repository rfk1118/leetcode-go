package main

// ListNode 基本数据结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if nil == head || head.Next == nil || head.Next.Next == nil {
		return false
	}
	slowIndex := head.Next
	fastIndex := head.Next.Next
	pos := -1
	for slowIndex != fastIndex {
		if fastIndex.Next == nil || fastIndex.Next.Next == nil {
			return false
		}
		slowIndex = slowIndex.Next
		fastIndex = fastIndex.Next.Next
	}
	fastIndex = head
	pos = 0
	for slowIndex != fastIndex {
		fastIndex = fastIndex.Next
		slowIndex = slowIndex.Next
		pos++
	}

	return slowIndex == fastIndex
}
