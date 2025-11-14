package main

import (
	"fmt"
	"net"
	"os"
	"io"
	"github.com/Drime648/coding-challenges/redis/internal/resp"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	aof, err := NewAof("redis_log.aof")
	if err != nil {
		fmt.Printf("Error with Creating Aof: %v\n", err)
		os.Exit(1)
	}
	defer aof.Close()



	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		go handleConnection(conn, aof)
	}
}


func handleConnection(conn net.Conn, aof *Aof) {
	defer conn.Close()

	respClient := resp.NewResp(conn)
	for {
		value, err := respClient.Read()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("error reading from client: %v\n", err)
			continue
		}
		// fmt.Println(value)

		if value.Typ != resp.TypeArray { //client needs to only send array
			fmt.Println("Invalid request, must be an array")
			continue
		}
		if len(value.Array) < 1 {
			fmt.Println("Invalid request, array must be of length >0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		callback, exists := Handlers[command]
		if !exists {
			msg := fmt.Sprintf("ERR unknown command '%s'", command)
			fmt.Println(msg)
			respClient.Write(resp.Value{Typ: resp.TypeError, Str: msg,})
			continue
		}

		responseVal, err := callback(value.Array[1:])
		if err != nil {
			fmt.Println(err)
			respClient.Write(resp.Value{Typ: resp.TypeError, Str: err.Error(),})
			continue
		}
		
		//command was successful here, so write to log
		//Only need to write SET or HSET because that will modify it.

		if command == "SET" || command == "HSET" {
			aof.Write(value)
		}


		// fmt.Println(responseVal)

		respClient.Write(responseVal)

		// respClient.Write(resp.Value{Typ: resp.TypeString, Str: "OK"})
	}
}
