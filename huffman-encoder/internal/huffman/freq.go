package huffman

import "io"

func CountFrequency(r io.Reader) (map[byte]int, error) {
	data := make([]byte, 4096)
	result := map[byte]int{}
	for {
		n, err := r.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return result, err
			}
		}
		for i := range n {
			b := data[i]
			result[b]++
		}
	}

	return result, nil
}
