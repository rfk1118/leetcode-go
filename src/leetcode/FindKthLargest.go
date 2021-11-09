package main

// PriorityQueue 需要存储的数据，需要大顶堆
type PriorityQueue struct {
	size int
	data []int
}

func newPriorityQueueWithItems(data []int) *PriorityQueue {
	queue := &PriorityQueue{}
	queue.size = len(data)
	queue.data = data
	queue.heapify()
	return queue
}

func (p *PriorityQueue) heapify() {
	lastFather := p.size/2 - 1
	for i := lastFather; i >= 0; i-- {
		p.shiftDown(i)
	}
}

func (p *PriorityQueue) shiftDown(index int) {
	half := p.size / 2
	for index < half {
		child := index*2 + 1
		maxValue := p.data[child]
		right := child + 1
		if right < p.size && p.data[right] > p.data[child] {
			maxValue = p.data[right]
			child = right
		}
		if maxValue < p.data[index] {
			break
		}
		p.data[child] = p.data[index]
		p.data[index] = maxValue
		index = child
	}
}

func (p *PriorityQueue) poll() int {
	r := p.data[0]
	p.size = p.size - 1
	p.data[0] = p.data[p.size]
	p.shiftDown(0)
	return r
}

//func main() {
//	intItems := make([]int, 10)
//	for i := 0; i < 10; i++ {
//		intItems[i] = i
//	}
//	fmt.Println(findKthLargest(intItems, 3))
//}

//
func findKthLargest(nums []int, k int) int {
	items := newPriorityQueueWithItems(nums)
	for i := 1; i < k; i++ {
		items.poll()
	}
	return items.poll()
}

func kthLargest(root *TreeNode, k int) int {
	t := make([]int, 0)
	var orderkthLargest func(node *TreeNode)
	orderkthLargest = func(node *TreeNode) {
		if nil == node {
			return
		}
		orderkthLargest(node.Right)
		t = append(t, node.Val)
		orderkthLargest(node.Left)
	}
	orderkthLargest(root)
	return t[k-1]
}

// [3,1,4,null,2]
func main() {
	node3 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node3.Left = node1
	node3.Right = node4
	node1.Left = node2

	kthLargest(node3, 1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
