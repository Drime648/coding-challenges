package huffman

import "io"

func buildPrefixTree(root *Node) map[byte]string {
	prefixTree := map[byte]string{}
	dfs(root, prefixTree, "")
	return prefixTree
}

func dfs(node *Node, prefixTree map[byte]string, currBitString string) {
	if node.left == nil && node.right == nil {
		prefixTree[node.data] = currBitString
	}

	if node.right != nil {
		dfs(node.right, prefixTree, currBitString+"1")
	}

	if node.left != nil {
		dfs(node.left, prefixTree, currBitString+"0")
	}
}

func encodeData(r io.Reader, w io.Writer) error {
	frequencyTable, err := CountFrequency(r)
	if err != nil {
		return err
	}
	root := buildEncoderTree(frequencyTable)
	prefixTree := buildPrefixTree(&root)
	data := make([]byte, 4096)
	for {
		n, err := r.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		for i := range n {
			b := data[i]
			bitString := prefixTree[b]

		}
	}

	return nil
}



func bitStringTobit(:)
