package huffman

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	last := old[n-1]

	*h = old[0 : n-1]

	return last
}
