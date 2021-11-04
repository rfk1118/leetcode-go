package main

func kthSmallest(root *TreeNode, k int) int {
	var heap []*TreeNode
	kthSmallestHelper(root, &heap)
	if k > len(heap) {
		return 0
	}
	return heap[k-1].Val

}

func kthSmallestHelper(root *TreeNode, heap *[]*TreeNode) {
	if root == nil {
		return
	}
	kthSmallestHelper(root.Left, heap)
	*heap = append(*heap, root)
	kthSmallestHelper(root.Right, heap)
}
