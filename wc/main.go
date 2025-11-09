/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/Drime648/coding-challenges/wc/internal/count"
	"os"
	"fmt"
	// "io"
)

type statsIndex int
const (
	NumBytesIndex statsIndex = 0
	NumLinesIndex statsIndex = 1
	NumWordsIndex statsIndex = 2
	NumCharsIndex statsIndex = 3
)

func main() {
	callbacks := make([]func(count.Stats) int, 4)
	input := os.Stdin
	var err error

	for _, cmd := range os.Args[1:] {
		switch cmd {
			case "-c":
				callbacks[NumBytesIndex] = func(s count.Stats) int {return s.NumBytes}
			
			case "-l":
				callbacks[NumLinesIndex] = func(s count.Stats) int {return s.NumLines}
			
			case "-w":
				callbacks[NumWordsIndex] = func(s count.Stats) int {return s.NumWords}
			
			case "-m":
				callbacks[NumCharsIndex] = func(s count.Stats) int {return s.NumChars}

			default:
				input, err = os.Open(cmd)
				if err != nil {
					fmt.Fprintf(os.Stderr,"Invalid Input: %v\n", err )
					os.Exit(1)
				}
		}
	}

	output := ""

	stats, err := count.CountData(input)
	if err != nil {
		fmt.Fprintf(os.Stderr,"error: %v\n", err )
		os.Exit(1)
	}

	for _, callback := range callbacks {
		if callback == nil {
			continue
		}
		count:= callback(stats)
		output = fmt.Sprintf("%s %d", output, count)

	}
	fmt.Printf("%s\n", output)
}
