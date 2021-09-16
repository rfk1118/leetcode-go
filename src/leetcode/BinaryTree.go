package main

import "fmt"

type BinaryTree struct {
	// key
	key   int
	value int

	// 左节点
	leftNode *BinaryTree
	// 右几点
	rightNode *BinaryTree
}

func newBinaryTree(key, value int, left, right *BinaryTree) *BinaryTree {
	return &BinaryTree{
		key:       key,
		value:     value,
		leftNode:  left,
		rightNode: right,
	}
}

func find(tree *BinaryTree, key int) *BinaryTree {
	if nil == tree {
		return nil
	}
	if tree.key == key {
		return tree
	} else if tree.key < key {
		return find(tree.rightNode, key)
	} else {
		return find(tree.leftNode, key)
	}
}

func findWithLoop(tree *BinaryTree, key int) *BinaryTree {
	if nil == tree {
		return nil
	}
	current := tree
	for current != nil && current.value != key {
		if current.key < key {
			current = current.rightNode
		} else {
			current = current.leftNode
		}
	}
	return current
}

func insert(head *BinaryTree, key, value int) *BinaryTree {
	if nil == head {
		head = newBinaryTree(key, value, nil, nil)
	}
	if head.key == key {
		head.value = value
	} else if head.key < key {
		head.rightNode = insert(head.rightNode, key, value)
	} else {
		head.leftNode = insert(head.leftNode, key, value)
	}
	return head
}

func insertWithLoop(head *BinaryTree, key, value int) *BinaryTree {
	if nil == head {
		head = newBinaryTree(key, value, nil, nil)
		return head
	}
	current := head
	for current.key != key {
		if current.key < key {
			if current.rightNode != nil {
				current = current.rightNode
			} else {
				current.rightNode = newBinaryTree(key, value, nil, nil)
				return current
			}
		} else if current.key > key {
			if current.leftNode != nil {
				current = current.leftNode
			} else {
				current.leftNode = newBinaryTree(key, value, nil, nil)
				return current
			}
		}
	}
	current.value = value
	return current
}

func preOrder(head *BinaryTree) {
	if nil == head {
		return
	} else {
		fmt.Printf("key:%d,value:%d\n", head.key, head.value)
	}
	preOrder(head.leftNode)
	preOrder(head.rightNode)
}

func preOrderWithLoop(head *BinaryTree) {
	// 如果头节点为空，直接返回
	if nil == head {
		return
	}
	// 创建栈
	stack := NewStack()
	// 推入头
	stack.Push(head)
	// 栈不为空
	for !stack.IsEmpty() {
		// 打印根
		pop := stack.Pop()
		// 拿出根，直接打印
		fmt.Printf("key:%d,value:%d\n", pop.key, pop.value)
		// 先推入右
		if nil != pop.rightNode {
			stack.Push(pop.rightNode)
		}
		// 在推入左
		if nil != pop.leftNode {
			stack.Push(pop.leftNode)
		}
	}
}

func inOrder(head *BinaryTree) {
	if nil == head {
		return
	}
	inOrder(head.leftNode)
	fmt.Printf("key:%d,value:%d\n", head.key, head.value)
	inOrder(head.rightNode)
}

func inOrderWithLoop(head *BinaryTree) {
	current := head
	// 创建一个辅助栈
	stack := NewStack()
	// 当前节点不为空或者栈内有数据
	for current != nil || !stack.IsEmpty() {
		if current != nil {
			// 先推入根节点
			stack.Push(current)
			// 如果根节点的左孩子不为空，在推入左孩子节点，这样出栈的时候就成了左右
			current = current.leftNode
		} else {
			// 推出一个节点
			current = stack.Pop()
			// 先左后根
			fmt.Printf("key:%d,value:%d\n", current.key, current.value)
			current = current.rightNode
		}
	}
}

func postOrder(head *BinaryTree) {
	if nil == head {
		return
	}
	postOrder(head.leftNode)
	postOrder(head.rightNode)
	fmt.Printf("key:%d,value:%d\n", head.key, head.value)
}

func postOrderWithLoop(head *BinaryTree) {
	// 如果头节点为空，直接返回
	if nil == head {
		return
	}
	stack := NewStack()
	// 出栈 左 右 根
	resultStack := NewStack()
	// 推入头
	stack.Push(head)
	// 栈不为空
	for !stack.IsEmpty() {
		// 先从栈中拿到根
		pop := stack.Pop()
		// 将根推到结果
		resultStack.Push(pop)
		// 在推入左
		if nil != pop.leftNode {
			stack.Push(pop.leftNode)
		}
		// 在推入右
		if nil != pop.rightNode {
			stack.Push(pop.rightNode)
		}
	}
	for !resultStack.IsEmpty() {
		pop := resultStack.Pop()
		fmt.Printf("key:%d,value:%d\n", pop.key, pop.value)
	}
}

func postOrderWithLoopOneStack(head *BinaryTree) {
	if nil == head {
		return
	}
	// 辅助栈
	stack := NewStack()
	// 头推入
	stack.Push(head)
	// 当前节点
	current := head
	for !stack.IsEmpty() {
		// 拿到了根节点但是当前先不推出
		peekNode := stack.peek()
		//  这里判断左节点不为空的时候才处理
		//  为什么这里peekNode.leftNode != current，因为可能上一次推出的是左孩子节点，这里刚回到父亲在推入不出问题了吗
		//  为什么peekNode.rightNode != current，因为这里推完左孩子，已经处理右孩子，而且处理完右孩子，所以当前等于右孩子，如果不判断的话就重复推入左孩子了
		if peekNode.leftNode != nil && peekNode.leftNode != current && peekNode.rightNode != current {
			// 推入左孩子节点
			stack.Push(peekNode.leftNode)
		} else if peekNode.rightNode != nil && peekNode.rightNode != current {
			// 推入右孩子节点
			stack.Push(peekNode.rightNode)
		} else {
			pop := stack.Pop()
			fmt.Printf("key:%d,value:%d\n", pop.key, pop.value)
			current = peekNode
		}
	}
}

func findMin(head *BinaryTree) *BinaryTree {
	if head == nil {
		return nil
	}
	if head.leftNode != nil {
		return findMin(head.leftNode)
	}
	return head
}

func findMinWithLoop(head *BinaryTree) *BinaryTree {
	if nil == head {
		return nil
	}
	current := head
	for current.leftNode != nil {
		current = current.leftNode
	}
	return current
}

func findMax(head *BinaryTree) *BinaryTree {
	if head == nil {
		return nil
	}
	if head.rightNode != nil {
		return findMax(head.rightNode)
	}
	return head
}

func findMaxWithLoop(head *BinaryTree) *BinaryTree {
	if nil == head {
		return nil
	}
	current := head
	for current.rightNode != nil {
		current = current.rightNode
	}
	return current
}

func main() {
	tree := insert(nil, 3, 3)
	insert(tree, 2, 2)
	insert(tree, 4, 4)
	insert(tree, 1, 1)
	insert(tree, 5, 5)
	fmt.Println("--------")
	postOrder(tree)
	fmt.Println("--------")
	postOrderWithLoop(tree)
	fmt.Println("--------")
	postOrderWithLoopOneStack(tree)
	fmt.Println("--------")
	//postOrder(tree)
	//fmt.Println("--------")
	//loop := findMin(tree)
	//fmt.Println(loop)
	//treeLoop := insertWithLoop(nil, 60, 60)
	//insert(treeLoop, 40, 40)
	//insert(treeLoop, 30, 30)
	//insert(treeLoop, 50, 50)
	//insert(treeLoop, 45, 45)
	//fmt.Println(treeLoop)
}

type Stack struct {
	e []*BinaryTree
}

func NewStack() *Stack {
	return &Stack{
		e: make([]*BinaryTree, 0),
	}
}
func (s *Stack) Len() int {
	return len(s.e)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(element *BinaryTree) bool {
	s.e = append(s.e, element)
	return true
}

func (s *Stack) Pop() *BinaryTree {
	if s.IsEmpty() {
		return nil
	}
	length := s.Len() - 1
	result := s.e[length]
	s.e = s.e[:length]
	return result
}

func (s *Stack) peek() *BinaryTree {
	if s.IsEmpty() {
		return nil
	}
	length := s.Len() - 1
	return s.e[length]
}

func (s *Stack) Show() {
	for _, element := range s.e {
		fmt.Println(element)
	}
}
