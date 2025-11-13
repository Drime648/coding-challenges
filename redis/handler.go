package main

import (
	// "fmt"
	// "net"
	// "os"
	// "io"
	"github.com/Drime648/coding-challenges/redis/internal/resp"
)


var Handlers = map[string]func([]resp.Value) resp.Value {
	"COMMAND": handleCommand,
	"PING": handlePing,
}

func handleCommand(args []resp.Value) resp.Value {
	return resp.Value{Typ: resp.TypeNull}
}


func handlePing(args []resp.Value) resp.Value {
	return resp.Value{Typ: resp.TypeString, Str: "PONG",}
}
