package huffman

import (
	"container/heap"
)

type Node struct {
	count int
	data  byte
	left  *Node
	right *Node
}

func buildHeap(freq map[byte]int) *NodeHeap {
	h := &NodeHeap{}
	heap.Init(h)
	for k, v := range freq {
		n := Node{count: v, data: k}
		heap.Push(h, n)
	}

	return h
}

func buildEncoderTree(freq map[byte]int) Node {
	nodeHeap := buildHeap(freq)

	for nodeHeap.Len() > 1 {
		n1 := heap.Pop(nodeHeap).(Node)
		n2 := heap.Pop(nodeHeap).(Node)

		sum := n1.count + n2.count
		newNode := Node{count: sum, left: &n1, right: &n2}
		heap.Push(nodeHeap, newNode)
	}

	return heap.Pop(nodeHeap).(Node)
}
