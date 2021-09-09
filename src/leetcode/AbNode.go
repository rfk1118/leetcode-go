package main

import "fmt"

type ABNode struct {
	next *ABNode
	v    int
}

// 使用快蛮步删除中间节点
func slowQuickRemoveMid(head *ABNode) *ABNode {
	if nil == head || head.next == nil {
		return head
	}
	slowCurrent := head
	fastCurrent := head.next.next

	for fastCurrent.next != nil && fastCurrent.next.next != nil {
		slowCurrent = slowCurrent.next
		fastCurrent = fastCurrent.next.next
	}

	slowCurrent.next = slowCurrent.next.next
	return head
}

func main() {
	node01 := &ABNode{nil, 1}
	node02 := &ABNode{nil, 3}
	node03 := &ABNode{nil, 5}
	node04 := &ABNode{nil, 7}
	node05 := &ABNode{nil, 9}
	node06 := &ABNode{nil, 11}
	node07 := &ABNode{nil, 12}
	node01.next = node02
	node02.next = node03
	node03.next = node04
	node04.next = node05
	node05.next = node06
	node06.next = node07

	node := slowQuickRemoveMid(node01)
	fmt.Println(node)
}
