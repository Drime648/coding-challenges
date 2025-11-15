package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Drime648/coding-challenges/json/internal/parser"
)

func main() {
	test := "{\"name\": \"dhruv\", \"age\": \"19\"}"
	x := strings.NewReader(test)
	rd := bufio.NewReader(x)

	output, err := parser.ParseObject(rd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)

	y := map[string]string{
		"hi": "bye",
	}
	fmt.Println(y)
}
