package huffman

func buildHeap(freq map[byte]int) *NodeHeap {
	heap := &NodeHeap{}
	for k, v := range freq {
		n := Node{count: v, data: k}
		heap.Push(n)
	}

	return heap
}

// build a heap first??
func buildEncoderTree(freq map[byte]int) Node {
	nodeHeap := buildHeap(freq)

	return Node{}
}
