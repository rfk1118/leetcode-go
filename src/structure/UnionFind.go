package main

type UnionFind struct {
	L []int
}

func newUnionFind(initSize int) *UnionFind {
	// 初始化数据结构
	uf := &UnionFind{
		L: make([]int, initSize),
	}
	// 自己的父亲指向自己
	for index := range uf.L {
		uf.L[index] = index
	}
	return uf
}

func (uf *UnionFind) find(item int) int {
	// 自己就是根节点
	if uf.L[item] == item {
		return item
	}
	// 路径压缩
	root := uf.find(uf.L[item])
	uf.L[item] = root
	return root
}

func (uf *UnionFind) union(p, q int) {
	uf.L[uf.find(p)] = uf.find(q)
}

func (uf *UnionFind) connection(p, q int) bool {
	return uf.find(p) == uf.find(q)
}
