package main

import (
	"fmt"
	// "net"
	// "os"
	// "io"
	"github.com/Drime648/coding-challenges/redis/internal/resp"
	"sync"
)


type Database struct {
	entries map[string]string
	mutex sync.RWMutex
}

var db = Database{}



var Handlers = map[string]func([]resp.Value) (resp.Value, error) {
	"COMMAND": handleCommand,
	"PING": handlePing,
	"SET": handleSet,
}

func handleCommand(args []resp.Value) (resp.Value, error) {
	return resp.Value{Typ: resp.TypeNull}, nil
}


func handlePing(args []resp.Value) (resp.Value, error) {
	return resp.Value{Typ: resp.TypeString, Str: "PONG",}, nil
}

func handleSet(args []resp.Value) (resp.Value, error) {
	if len(args) != 2{
		return resp.Value{}, fmt.Errorf("Invalid Number of arguments")
	}
	if args[0].Typ != resp.TypeBulk || args[1].Typ != resp.TypeBulk {
		return resp.Value{}, fmt.Errorf("Invalid argument types. Must be Bulk Strings")
	}

	key := args[0].Bulk
	val := args[1].Bulk

	db.mutex.Lock()
	db.entries[key] = val
	db.mutex.Unlock()

	return resp.Value{Typ:resp.TypeString, Str: "OK"}, nil

}
