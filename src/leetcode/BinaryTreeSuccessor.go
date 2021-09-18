package main

type BinaryTreeSuccessor struct {
	key, value int
	leftNode   *BinaryTreeSuccessor
	rightNode  *BinaryTreeSuccessor
	parent     *BinaryTreeSuccessor
}

func findBinaryTreeSuccessor(head *BinaryTreeSuccessor) *BinaryTreeSuccessor {
	// 当前节点的右节点不为空情况下
	if head.rightNode != nil {
		return findMinBinaryTreeSuccessor(head)
	}

	// 第二种情况，当前节点就是节点B的左节点，所以节点B就是当前节点后继节点
	current := head
	currentParent := head.parent
	// 当前节点就是第情况2，直接就不循环了
	for currentParent.leftNode != current {
		current = currentParent
		currentParent = currentParent.parent
	}
	return currentParent
}

// 右节点不为空的情况下，右子树不为空
func findMinBinaryTreeSuccessor(head *BinaryTreeSuccessor) *BinaryTreeSuccessor {
	for head.leftNode != nil {
		head = head.leftNode
	}
	return head
}
