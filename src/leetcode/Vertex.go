package main

import "fmt"

type Vertex struct {
	// 这里是数据，后面可以转换成对象
	label rune
	// 是否访问过，在遍历的时候很有用
	wasVisited bool
}

func newVertex(label rune) *Vertex {
	return &Vertex{label: label, wasVisited: false}
}

const MaxVertex int = 20
const LINK int = 1

type Graph struct {
	vertexList []*Vertex
	adjMat     [][]int
	nVertex    int
}

func newGraph() *Graph {
	return &Graph{
		vertexList: make([]*Vertex, MaxVertex),
		adjMat:     make([][]int, MaxVertex, MaxVertex),
		nVertex:    0,
	}
}

func (g *Graph) addVertex(label rune) {
	g.nVertex++
	g.vertexList[g.nVertex] = newVertex(label)
}

func (g *Graph) addEdge(start, end int) {
	g.adjMat[start][end] = LINK
	g.adjMat[end][start] = LINK
}

func (g *Graph) getAdjUnvisitedVertex(v int) int {
	for j := 0; j < g.nVertex; j++ {
		// 连接并且没有访问过
		if g.adjMat[v][j] == LINK && !g.vertexList[v].wasVisited {
			return j
		}
	}
	return -1
}

func (g *Graph) displayVertex(v int) {
	fmt.Println(g.vertexList[v].label)
}

type StackX struct {
	st       [MaxVertex]int
	topIndex int
}

func newStackX() *StackX {
	return &StackX{
		topIndex: -1,
	}
}

func (s *StackX) push(item int) {
	s.topIndex = s.topIndex + 1
	s.st[s.topIndex] = item
}

func (s *StackX) pop() int {
	r := s.st[s.topIndex]
	s.topIndex = s.topIndex - 1
	return r
}

func (s *StackX) peek() int {
	r := s.st[s.topIndex]
	return r
}

func (s *StackX) isEmpty() bool {
	return s.topIndex == -1
}

func (g *Graph) dfs() {
	stackX := newStackX()
	g.vertexList[0].wasVisited = true
	g.displayVertex(0)
	stackX.push(0)
	for !stackX.isEmpty() {
		currentNode := g.getAdjUnvisitedVertex(stackX.peek())
		if currentNode == -1 {
			stackX.pop()
		} else {
			g.vertexList[currentNode].wasVisited = true
			g.displayVertex(currentNode)
			stackX.push(currentNode)
		}
	}
	for _, itemValue := range g.vertexList {
		itemValue.wasVisited = false
	}
}

type QueueX struct {
	q     [MaxVertex]int
	front int
	rear  int
}

func newQueueX() *QueueX {
	return &QueueX{
		front: 0,
		rear:  -1,
	}
}

func (r *QueueX) insert(item int) {
	if r.rear == len(r.q)-1 {
		r.rear = -1
	}
	r.q[r.rear] = item
}

func (r *QueueX) remove() int {
	r.front++
	temp := r.q[r.front]
	if r.front == len(r.q) {
		r.front = 0
	}
	return temp
}

func (r *QueueX) isEmpty() bool {
	return r.rear+1 == r.front || (r.front+len(r.q)-1 == r.rear)
}

func (g *Graph) bfs() {
	queueX := newQueueX()
	g.vertexList[0].wasVisited = true
	g.displayVertex(0)
	queueX.insert(0)
	for !queueX.isEmpty() {
		remove := queueX.remove()
		vertex := g.getAdjUnvisitedVertex(remove)
		for vertex != -1 {
			g.vertexList[vertex].wasVisited = true
			g.displayVertex(vertex)
			queueX.insert(vertex)
			vertex = g.getAdjUnvisitedVertex(vertex)
		}
	}
	for _, itemValue := range g.vertexList {
		itemValue.wasVisited = false
	}
}

func (g *Graph) mst() {
	stackX := newStackX()
	g.vertexList[0].wasVisited = true
	stackX.push(0)
	for !stackX.isEmpty() {
		// 能找到其他连接，继续向下走
		from := stackX.peek()
		to := g.getAdjUnvisitedVertex(from)
		if to == -1 {
			stackX.pop()
		} else {
			g.vertexList[to].wasVisited = true
			stackX.push(to)
			g.displayVertex(from)
			g.displayVertex(to)
		}
	}
	for _, itemValue := range g.vertexList {
		itemValue.wasVisited = false
	}
}

func (g *Graph) topo() {
	var st [MaxVertex]rune
	vertex := g.nVertex
	for vertex > 0 {
		currentNode := g.noSuccessors()
		if currentNode == vertex-1 {
			fmt.Println("error graph have cycles")
			return
		}
		st[g.nVertex-1] = g.vertexList[currentNode].label
		g.deleteVertex(currentNode)
	}
	fmt.Println("sort order")
	for _, r := range st {
		fmt.Println(r)
	}
}

// 查找后继续节点
func (g *Graph) noSuccessors() int {
	var isEdge bool
	vertex := g.nVertex
	for row := 0; row < vertex; row++ {
		isEdge = false
		for col := 0; col < vertex; col++ {
			if g.adjMat[row][col] > 0 {
				isEdge = true
				break
			}
		}
		if !isEdge {
			return row
		}
	}
	return -1
}

func (g *Graph) deleteVertex(delVel int) {
	if delVel != g.nVertex-1 {
		for j := delVel; j < g.nVertex-1; j++ {
			g.vertexList[j] = g.vertexList[j+1]
		}
		for row := delVel; row < g.nVertex-1; row++ {
			g.moveRowUp(row, g.nVertex)
		}
		for col := delVel; col < g.nVertex-1; col++ {
			g.moveColLeft(col, g.nVertex)
		}
	}
	g.nVertex = g.nVertex - 1
}

func (g *Graph) moveRowUp(row int, length int) {
	for col := 0; col < length; col++ {
		g.adjMat[row][col] = g.adjMat[row+1][col]
	}
}

func (g *Graph) moveColLeft(col int, length int) {
	for row := 0; row < length; row++ {
		g.adjMat[row][col] = g.adjMat[row][col+1]
	}
}
