package main

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if nil == headA || nil == headB {
		return nil
	}
	currentHeadA := headA
	currentHeadALength := 0
	currentHeadAPre := headA
	for currentHeadA != nil {
		currentHeadAPre = currentHeadA
		currentHeadA = currentHeadA.Next
		currentHeadALength++
	}
	currentHeadBPre := headA
	currentHeadB := headB
	currentHeadBLength := 0
	for currentHeadB != nil {
		currentHeadBPre = currentHeadB
		currentHeadB = currentHeadB.Next
		currentHeadBLength++
	}
	// 最后一个节点是否为同一个节点
	if currentHeadAPre != currentHeadBPre {
		return nil
	}
	currentHeadA = headA
	currentHeadB = headB
	if currentHeadALength > currentHeadBLength {
		step := currentHeadALength - currentHeadBLength
		for i := 0; i < step; i++ {
			currentHeadA = currentHeadA.Next
		}
	} else {
		step := currentHeadBLength - currentHeadALength
		for i := 0; i < step; i++ {
			currentHeadB = currentHeadB.Next
		}
	}
	for currentHeadA != currentHeadB {
		currentHeadA = currentHeadA.Next
		currentHeadB = currentHeadB.Next
	}
	return currentHeadA

}
