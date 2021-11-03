package main

// map里面保存数据
// 我可能对这个题有误解
// 1. 值可能为负数，所以不能直接当index,因为union不能存在index为负数的情况
// 2. 连接不一定为连续最大值，比如[0,2]和[1,9]这里连接比较多，但是只有1～9属于正常情况
type indexWrap struct {
	// 在使用Go必须用个对象保存这个东西要不map[int]bool找不到index
	index int
}
type let123 struct {
	parent []int
}

func newLet123(nums []int) *let123 {
	l := &let123{
		parent: make([]int, len(nums)),
	}
	for i := 0; i < len(l.parent); i++ {
		l.parent[i] = i
	}
	return l
}

func (l *let123) find(index int) int {
	if index == l.parent[index] {
		return index
	}
	i := l.find(l.parent[index])
	l.parent[index] = i
	return i
}

func (l *let123) union(p, q int) {
	rootP := l.find(p)
	rootQ := l.find(q)
	if rootQ == rootP {
		return
	}
	// q的值会大，让p的父亲为q
	l.parent[rootP] = l.parent[rootQ]
}

func longestConsecutive(nums []int) int {
	result := 0
	if nil == nums || len(nums) == 0 {
		return result
	}
	// map使用int bool 类型的话，int可能值会为负数，在union中使用index会错误
	// 如果使用int int,对于查找不到的值，会返回值0，与角标为0的值在解释上会有冲突
	m := make(map[int]*indexWrap, len(nums))
	for index, num := range nums {
		wrap := m[num]
		// 不会覆盖值
		if nil != wrap {
			continue
		}
		m[num] = &indexWrap{
			index: index,
		}
	}
	union := newLet123(nums)
	for index, wrap := range m {
		// 这里找到的是大值
		indexWrap := m[index+1]
		if nil != indexWrap {
			// 连接大值
			union.union(wrap.index, indexWrap.index)
		}
	}

	for i := 0; i < len(nums); i++ {
		root := union.find(i)
		max := nums[root] - nums[i]
		if max > result {
			result = max
		}
	}
	return result + 1
}
