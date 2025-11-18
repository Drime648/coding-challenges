package huffman

type Node struct {
	count int
	data  byte
	left  *Node
	right *Node
}
