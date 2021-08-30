package main

import "fmt"

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

func main() {
	intItems := make([]int, 10)
	for i := 0; i < 10; i++ {
		intItems[i] = i
	}
	fmt.Println(findKthLargest(intItems, 3))
}

//
func findKthLargest(nums []int, k int) int {
	items := newPriorityQueueWithItems(nums)
	for i := 1; i < k; i++ {
		items.poll()
	}
	return items.poll()
}
