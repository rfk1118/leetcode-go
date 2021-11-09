package main

func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool
	var dfs func(node *TreeNode, p *TreeNode, x int, y int, l int)
	dfs = func(node *TreeNode, p *TreeNode, x int, y int, l int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent, xDepth, xFound = p, l, true
		}
		if node.Val == y {
			yParent, yDepth, yFound = p, l, true
		}
		if xFound && yFound {
			return
		}
		dfs(node.Left, node, x, y, l+1)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, x, y, l+1)
	}
	dfs(root, nil, x, y, 1)
	return xParent != yParent && xDepth == yDepth
}

func isCousinsBfs(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool
	// only handler
	update := func(node, parent *TreeNode, depth int) {
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		}
		if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
	}

	type pair struct {
		node  *TreeNode
		depth int
	}
	q := []pair{{root, 0}}
	update(root, nil, 0)
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		if pop.node.Left != nil {
			q = append(q, pair{pop.node.Left, pop.depth + 1})
			update(pop.node.Left, pop.node, pop.depth+1)
		}
		if pop.node.Right != nil {
			q = append(q, pair{pop.node.Right, pop.depth + 1})
			update(pop.node.Right, pop.node, pop.depth+1)
		}
	}
	return xParent != yParent && xDepth == yDepth
}
