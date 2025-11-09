/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/Drime648/coding-challenges/wc/internal/count"
	"os"
	"fmt"
)

type commands struct {
	bytesCount bool
	wordsCount bool
	linesCount bool
	charsCount bool
}

func main() {

	c := commands {
		bytesCount: false,
		wordsCount: false,
		linesCount: false,
		charsCount: false,
	}

	input := os.Stdin
	var err error

	for _, cmd := range os.Args[1:] {
		switch cmd {
			case "-c":
				c.bytesCount = true
			
			case "-l":
				c.linesCount = true
			
			case "-w":
				c.wordsCount = true
			
			case "-m":
				c.charsCount = true
			default:
				input, err = os.Open(cmd)
				if err != nil {
					fmt.Fprintf(os.Stderr,"Invalid Input: %v\n", err )
					os.Exit(1)
				}
		}
	}

	output := ""
	if c.bytesCount {
		bytesCount, err := count.CountBytes(input)
		if err != nil {
			fmt.Fprintf(os.Stderr,"error with counting bytes: %v\n", err )
			os.Exit(1)
		}
		output = fmt.Sprintf("%s %d", output, bytesCount)
	}

	if c.linesCount {
		linesCount, err := count.CountLines(input)
		if err != nil {
			fmt.Fprintf(os.Stderr,"error with counting lines: %v\n", err )
			os.Exit(1)
		}
		output = fmt.Sprintf("%s %d", output, linesCount)
	}

	if c.wordsCount {
		wordsCount, err := count.CountWords(input)
		if err != nil {
			fmt.Fprintf(os.Stderr,"error with counting words: %v\n", err )
			os.Exit(1)
		}
		output = fmt.Sprintf("%s %d", output, wordsCount)
	}
	fmt.Printf("%s\n", output)
}
