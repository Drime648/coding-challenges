package main

import (
	"fmt"
	"net"
	"os"
	"io"
	"github.com/Drime648/coding-challenges/redis/internal/resp"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		go handleConnection(conn)
	}
}


func handleConnection(conn net.Conn) {
	defer conn.Close()

	respClient := resp.NewResp(conn)
	for {
		value, err := respClient.Read()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("error reading from client: %v", err)
			continue
		}
		fmt.Println(value)

		respClient.Write(resp.Value{Typ: resp.TypeString, Str: "OK"})
	}
}
