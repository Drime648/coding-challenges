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

	flagFuncs := map[statsIndex]func(count.Stats) int {
		NumBytesIndex: func(s count.Stats) int {return s.NumBytes},
		NumLinesIndex: func(s count.Stats) int {return s.NumLines},
		NumWordsIndex: func(s count.Stats) int {return s.NumWords},
		NumCharsIndex: func(s count.Stats) int {return s.NumChars},
	}
	
	flagFound := false
	outputName := ""

	for _, cmd := range os.Args[1:] {
		switch cmd {
			case "-c":
				callbacks[NumBytesIndex] = flagFuncs[NumBytesIndex]
				flagFound = true
			
			case "-l":
				callbacks[NumLinesIndex] = flagFuncs[NumLinesIndex]
				flagFound = true
			
			case "-w":
				callbacks[NumWordsIndex] = flagFuncs[NumWordsIndex]
				flagFound = true
			
			case "-m":
				callbacks[NumCharsIndex] = flagFuncs[NumCharsIndex]
				flagFound = true

			default:
				input, err = os.Open(cmd)
				if err != nil {
					fmt.Fprintf(os.Stderr,"Invalid Input: %v\n", err )
					os.Exit(1)
				}
				outputName = cmd
		}
	}
	if !flagFound { //no flags means I do everything.
		callbacks[NumBytesIndex] = flagFuncs[NumBytesIndex]
		callbacks[NumLinesIndex] = flagFuncs[NumLinesIndex]
		callbacks[NumWordsIndex] = flagFuncs[NumWordsIndex]
		callbacks[NumCharsIndex] = flagFuncs[NumCharsIndex]
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
	fmt.Printf("%s %s\n", output, outputName)
}
