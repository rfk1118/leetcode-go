package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

func findSameNode(head01, head02 *Node) []*Node {
	if nil == head01 && nil == head02 {
		return nil
	}
	if nil == head01 && nil != head02 {
		return nil
	}
	if nil == head02 && nil != head01 {
		return nil
	}
	result := make([]*Node, 0)
	current01 := head01
	current02 := head02
	for current01 != nil && current02 != nil {
		if current01.value == current02.value {
			result = append(result, current01)
			current02 = current02.next
			current01 = current01.next
		} else if current01.value < current02.value {
			current01 = current01.next
		} else if current02.value < current01.value {
			current02 = current02.next
		}
	}
	return result
}

func main() {
	node01 := &Node{nil, 1}
	node02 := &Node{nil, 3}
	node03 := &Node{nil, 5}
	node04 := &Node{nil, 7}
	node05 := &Node{nil, 9}
	node06 := &Node{nil, 11}
	node01.next = node02
	node02.next = node03
	node03.next = node04
	node04.next = node05
	node05.next = node06

	node011 := &Node{nil, 2}
	node033 := &Node{nil, 5}
	node055 := &Node{nil, 7}
	node066 := &Node{nil, 8}
	node011.next = node033
	node033.next = node055
	node055.next = node066

	allNode := findSameNode(node01, node011)
	for _, node := range allNode {
		fmt.Println(node.value)
	}

}
