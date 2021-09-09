package main

import "fmt"

type RNode struct {
	next  *RNode
	value int
}

// 删除倒数第k个节点
func removeLastKthNode(head *RNode, kth int) *RNode {
	// 如果头为空直接返回，如果kth小于1的话直接返回
	if nil == head || kth < 1 {
		return nil
	}

	currentNode := head
	k := kth
	for currentNode != nil {
		k--
		currentNode = currentNode.next
	}

	//  如果大于0
	if k > 0 {
		return head
	}

	if k == 0 {
		return head.next
	}

	currentNode = head
	k++
	for k < 0 {
		currentNode = currentNode.next
		k++
	}

	currentNode.next = currentNode.next.next
	return head
}

func main() {
	node01 := &RNode{nil, 1}
	node02 := &RNode{nil, 3}
	node03 := &RNode{nil, 5}
	node04 := &RNode{nil, 7}
	node05 := &RNode{nil, 9}
	node06 := &RNode{nil, 11}
	node01.next = node02
	node02.next = node03
	node03.next = node04
	node04.next = node05
	node05.next = node06

	node := removeLastKthNode(node01, 3)
	fmt.Println(node)
}
