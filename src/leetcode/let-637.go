package main

var m map[int]*averageOfLevel

type averageOfLevel struct {
	value int
	count int
}

func averageOfLevels(root *TreeNode) (r []float64) {
	m = make(map[int]*averageOfLevel, 0)
	averageOfLevelsHelper(root, 1)
	for i := 1; i <= len(m); i++ {
		a := m[i]
		r = append(
			r,
			float64(a.value)/float64(a.count),
		)
	}
	return r
}

func averageOfLevelsHelper(root *TreeNode, lev int) {
	if nil == root {
		return
	}
	v := m[lev]
	if v == nil {
		level := &averageOfLevel{
			value: root.Val,
			count: 1,
		}
		m[lev] = level
	} else {
		m[lev] = &averageOfLevel{
			value: v.value + root.Val,
			count: v.count + 1,
		}
	}
	averageOfLevelsHelper(root.Left, lev+1)
	averageOfLevelsHelper(root.Right, lev+1)
}

// 这个就是我想要的
func averageOfLevelss(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		a := float64(sum) / float64(size)
		res = append(res, a)
		queue = queue[size:]
	}
	return res
}
