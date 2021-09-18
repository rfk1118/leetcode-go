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
	tree := insert(nil, 10, 10)
	insert(tree, 38, 38)
	insert(tree, 26, 26)
	insert(tree, 72, 72)
	insert(tree, 55, 55)
	//insert(tree, 90, 90)
	//insert(tree, 41, 41)
	//insert(tree, 43, 43)
	//insert(tree, 60, 60)
	//insert(tree, 78, 78)
	//insert(tree, 92, 92)
	//insert(tree, 74, 74)
	//fmt.Println("--------")
	//deleteKey(tree, 38)
	//fmt.Println("--------")
	//fmt.Println(tree)
	fmt.Println(binaryTreeSize(tree))
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

func deleteKeyWithLoop(head *BinaryTree, key int) bool {
	current := head
	parent := head
	isLeftChild := true
	// 这里是查找key
	for key != current.key {
		parent = current
		if current.key > key {
			isLeftChild = true
			current = current.leftNode
		} else {
			isLeftChild = false
			current = current.rightNode
		}
		if current == nil {
			return false
		}
	}
	// 当前只删除了叶子节点
	if current.leftNode == nil && current.rightNode == nil {
		if current == head {
			head = nil
		} else if isLeftChild {
			parent.leftNode = nil
		} else {
			parent.rightNode = nil
		}
	} else if current.rightNode == nil {
		if current == head {
			head = current.leftNode
		} else if isLeftChild {
			parent.leftNode = current.leftNode
		} else {
			parent.rightNode = current.leftNode
		}
	} else if current.leftNode == nil {
		if current == head {
			head = current.rightNode
		} else if isLeftChild {
			parent.leftNode = current.rightNode
		} else {
			parent.rightNode = current.rightNode
		}
	} else {
		successor := findSuccessor(current)
		// 删除的是根节点
		if current == head {
			head = successor
		} else if isLeftChild {
			parent.leftNode = successor
		} else {
			// 删除节点的父亲的右节点设置为后继
			parent.rightNode = successor
		}
		// 后继节点的左子树为删除节点的左子树
		successor.leftNode = current.leftNode
	}
	return true
}

func findSuccessor(head *BinaryTree) *BinaryTree {
	// 开始是后继节点父亲为要删除节点
	successorParent := head
	// 开始是后继节点为要删除节点
	successor := head
	//  拿到右节点
	current := head.rightNode
	// 这里要往左边走。如果左边不为空，一直走
	for current != nil {
		// 每向下一步，每次保存父亲节点
		successorParent = successor
		// 后继节点为当前节点
		successor = current
		// 继续向左走，这里可能为空，为空的话，successor = current，这里已经把 current设置成后继
		current = current.leftNode
	}
	// 因为后继节点已经没有左子树了
	if successor != head.rightNode {
		// 让后继节点的父亲的的左孩子等于右继续节点的左孩子
		successorParent.leftNode = successor.rightNode
		//  把要删除的节点的右边给后继节点
		successor.rightNode = head.rightNode
	}
	return successor
}

func deleteKey(head *BinaryTree, key int) *BinaryTree {
	// 没有找到
	if head == nil {
		return nil
	}
	if head.key > key {
		head.leftNode = deleteKey(head.leftNode, key)
	} else if head.key < key {
		head.rightNode = deleteKey(head.rightNode, key)
	} else {
		//  删除节点没有孩子,让其父亲删除该节点
		if head.leftNode == nil && head.rightNode == nil {
			return nil
		} else if head.leftNode == nil {
			return head.rightNode
		} else if head.rightNode == nil {
			return head.leftNode
		} else {
			// 从当前节点的左边开始找
			successorWrap := getSuccessor(newBinaryTreeWrap(head.rightNode, head))
			if successorWrap.successor != head.rightNode {
				successorWrap.successorParent.leftNode = successorWrap.successor.rightNode
				successorWrap.successor.rightNode = head.rightNode
			}
			// 后继节点为左节点为要删除的节点
			successorWrap.successor.leftNode = head.leftNode
			head = successorWrap.successor
		}
	}
	// 这里返回的节点让其父亲进行连接
	return head
}

func getSuccessor(head *BinaryTreeWrap) *BinaryTreeWrap {
	if head.successor.leftNode != nil {
		return getSuccessor(newBinaryTreeWrap(head.successor.leftNode, head.successor))
	} else {
		return head
	}
}

type BinaryTreeWrap struct {
	// 后继续节点
	successor       *BinaryTree
	successorParent *BinaryTree
}

func newBinaryTreeWrap(successor, successorParent *BinaryTree) *BinaryTreeWrap {
	return &BinaryTreeWrap{
		successor:       successor,
		successorParent: successorParent,
	}
}

func binaryTreeSize(tree *BinaryTree) int {
	// 如果子树为空返回0
	if tree == nil {
		return 0
	}
	// 递归计算左子树
	lSize := binaryTreeSize(tree.leftNode)
	// 递归计算右子树
	rSize := binaryTreeSize(tree.rightNode)
	// 加上当前节点
	return rSize + lSize + 1
}

func deleteKeyFour(head *BinaryTree, key int) *BinaryTree {
	// 没找到元素，不进行处理
	if nil == head {
		return nil
	}
	if head.key < key {
		// 如果要删除根节点的右节点，继续在右子树上查找
		head.rightNode = deleteKeyFour(head.rightNode, key)
	} else if head.key > key {
		// 如果要删除根节点的左节点，继续在左子树上查找
		head.leftNode = deleteKeyFour(head.leftNode, key)
	} else {
		// 要删除节点
		// 两个都为空
		if head.leftNode == nil && head.rightNode == nil {
			return nil
		}
		// 如果左节点为空，返回右节点
		if head.leftNode == nil {
			return head.rightNode
		}
		// 如果右节点为空，返回左节点
		if head.rightNode == nil {
			return head.leftNode
		}
		// 先查找到后继节点，也就是当前子树的右子树的最小节点
		successor := findMin(head.rightNode)
		// 因为后继节点可能是右子树的根节点，所以这时候的右节点就编程了原来的head.rightNode.rightNode
		// 1.如果后继节点为删除节点的右子节点，这时候successor.rightNode = head.rightNode.rightNode
		// 2.如果后继节点是删除节点的右子左孙节点，这时候，递归会把后继节点的右孩子连接到后继节点父亲的左子树上，这时候successor.rightNode = head.rightNode
		successor.rightNode = deleteMin(head.rightNode)
		// 后继续节点的左节点是删除节点的左节点
		successor.leftNode = head.leftNode
	}
	return head
}

func deleteMin(b *BinaryTree) *BinaryTree {
	// 因为左边没节点了，当前节点就是最小节点
	if nil == b.leftNode {
		return b.rightNode
	}
	// 还有小节点继续处理
	b.leftNode = deleteMin(b.leftNode)
	return b
}

func floor(b *BinaryTree, key int) *BinaryTree {
	// 如果找不到当前节点，直接返回空
	if b == nil {
		return nil
	}
	// 如果在左子树一定在
	if key < b.key {
		return floor(b.leftNode, key)
	}
	// 查询右子树
	if b.key < key {
		tree := floor(b.rightNode, key)
		if nil == tree {
			// 返回为空的话就说明右节点没有小于等于key的节点
			return b
		} else {
			return tree
		}
	}
	return b
}

func ceiling(b *BinaryTree, key int) *BinaryTree {
	// 子树为空，返回空
	if b == nil {
		return nil
	}
	//  key与子树的根节点相同
	if b.key == key {
		return b
	}
	// 如果在右子树上，一定存在
	if b.key < key {
		return ceiling(b.rightNode, key)
	}
	// 左子树不一定存在，因为没有大于等于key的节点的话，根节点就是结果
	r := ceiling(b.leftNode, key)
	if nil != r {
		return r
	}
	return b
}

func selection(b *BinaryTree, k int) *BinaryTree {
	size := binaryTreeSize(b.leftNode)
	if size > k {
		// 如果在左边，一直向左边缩小，每次数量会减少一半
		return selection(b.leftNode, k)
	} else if size < k {
		// 从右子树的左子树
		return selection(b.rightNode, k-size-1)
	} else {
		// 如果正好相等，则这个就是结果
		return b
	}
}

func rank(b *BinaryTree, key int) int {
	// 如果为空，返回0
	if b == nil {
		return 0
	}
	// 如果这个key相等，返回左子树数量
	if b.key == key {
		return binaryTreeSize(b.leftNode)
	} else if b.key > key {
		// 如果key小于子树的根，继续左子树查找
		return rank(b.leftNode, key)
	} else {
		// 如果在右子树，先加上所有左子树数量+根节点
		// 从右树上继续找左节点数量
		return 1 + binaryTreeSize(b.leftNode) + rank(b.rightNode, key)
	}
}

