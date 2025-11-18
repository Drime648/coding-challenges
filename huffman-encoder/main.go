package main

import (
	"fmt"
	"os"

	"github.com/Drime648/coding-challenges/huffman-encoder/internal/huffman"
)

func main() {
	f, _ := os.Open("test.txt")
	table, _ := huffman.CountFrequency(f)

	for k, v := range table {
		fmt.Println(string(k), ": ", v)
	}
}
