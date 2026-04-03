package main

import (
	"bufio"
	"log"
	"os"

	"lsp/rpc"
)

func main() {
	logger := getLogger("/Users/dhruv/coding/coding-challenges/lsp/log.out")
	logger.Println("Started LSP")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg, logger)
	}
}

func handleMessage(msg any, logger *log.Logger) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0o666)
	if err != nil {
		panic("bad file!")
	}

	return log.New(logfile, "[lsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
