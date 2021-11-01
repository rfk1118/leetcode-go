package main

type UF struct {
	parent, rank []int
	nComponents  int
}

func newUF(initSize int) *UF {
	// 初始化数据结构
	uf := &UF{
		parent:      make([]int, initSize),
		rank:        make([]int, initSize),
		nComponents: initSize,
	}
	// 子树等级初始化
	for index := range uf.rank {
		uf.rank[index] = 0
	}
	// 自己的父亲指向自己
	for index := range uf.parent {
		uf.rank[index] = index
	}
	return uf
}

func (uf *UF) find(item int) int {
	for item != uf.parent[item] {
		// 路径压缩
		uf.parent[item] = uf.parent[uf.parent[item]]
		// 向上走
		item = uf.parent[item]
	}
	return item
}

func (uf *UF) merge(p, q int) {
	// 查找p的根
	rootP := uf.find(p)
	// 查找q的根
	rootQ := uf.find(q)
	// 已经在一个集合里面了
	if rootQ == rootP {
		return
	}
	subTreeP := uf.rank[rootP]
	subTreeQ := uf.rank[rootQ]
	if subTreeP < subTreeQ {
		uf.parent[rootP] = rootQ
	} else if subTreeP > subTreeQ {
		uf.parent[rootQ] = rootP
	} else {
		// 子树等级相同，使用那个都行
		// uf.parent[rootP]  = rootQ
		// uf.parent[rootQ]  = rootP
		uf.parent[rootQ] = rootP
		uf.rank[rootP]++
	}
	// 集合个数减少一个
	uf.nComponents--
}
