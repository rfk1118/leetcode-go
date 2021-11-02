package main

type let200 struct {
	// 岛屿的数量，也就是集合的数量
	nComponent int
	// 这里保存根问题
	parent     []int
	// 等级
	rank       []int
}

func newLet200(grid [][]byte) *let200 {
	// 横向
	m := len(grid)
	// 纵
	n := len(grid[0])
	// 并查集
	l := &let200{
		nComponent: 0,
		parent:     make([]int, m*n),
		rank:       make([]int, m*n),
	}
	// 数组边界问题
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 创建所有集合
			if grid[i][j] == '1' {
				// 自己的父亲指向自己
				l.parent[i*n+j] = i*n + j
				l.nComponent++
			}
			// 初始化等级
			l.rank[i*n+j] = 0
		}
	}
	return l
}

func (l *let200) find(i int) int {
	if i == l.parent[i] {
		return i
	}

	root := l.find(l.parent[i])
	// 路径压缩
	l.parent[i] = root
	return root
}

func (l *let200) union(p, q int) {
	rootP := l.find(p)
	rootQ := l.find(q)
	// 已经连接过了
	if rootQ == rootP {
		return
	}
	// 等级小的连接到大的上去
	if l.rank[rootQ] > l.rank[rootP] {
		l.parent[rootP] = rootQ
	} else if l.rank[rootP] > l.rank[rootQ] {
		l.parent[rootQ] = rootP
	} else {
		l.parent[rootQ] = rootP
		l.rank[rootP]++
	}
	// 集合个数减少一个
	l.nComponent--
}

func numIslands(grid [][]byte) int {
	// 基本判断
	if nil == grid || len(grid) == 0 {
		return 0
	}
	// x*y
	m := len(grid)
	n := len(grid[0])
	// 创建并查集
	unionFind := newLet200(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					unionFind.union(i*n+j, (i-1)*n+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					unionFind.union(i*n+j, (i+1)*n+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					unionFind.union(i*n+j, i*n+j-1)
				}
				if j+1 < n && grid[i][j+1] == '1' {
					unionFind.union(i*n+j, i*n+j+1)
				}
			}
		}
	}
	return unionFind.nComponent
}
