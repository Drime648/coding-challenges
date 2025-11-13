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

var db = Database{entries: map[string]string{}}

type HashDatabase struct {
	entries map[string]map[string]string

	mutex sync.RWMutex
}

var hashDB = HashDatabase{entries: map[string]map[string]string{},}


var Handlers = map[string]func([]resp.Value) (resp.Value, error) {
	"COMMAND": handleCommand,
	"PING": handlePing,
	"SET": handleSet,
	"GET": handleGet,
	"HSET": handleHSet,
	"HGET": handleHGet,
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

func handleGet(args []resp.Value) (resp.Value, error) {
	if len(args) != 1{
		return resp.Value{}, fmt.Errorf("Invalid Number of arguments")
	}
	if args[0].Typ != resp.TypeBulk {
		return resp.Value{}, fmt.Errorf("Invalid argument types. Must be Bulk Strings")
	}

	key := args[0].Bulk

	db.mutex.RLock()
	val, exists := db.entries[key]
	db.mutex.RUnlock()
	
	if !exists {
		return resp.Value{Typ:resp.TypeNull}, nil
	}


	return resp.Value{Typ:resp.TypeBulk, Bulk: val}, nil
}

func handleHSet(args []resp.Value) (resp.Value, error) {
	if len(args) != 3 {
		return resp.Value{}, fmt.Errorf("Invalid Number of arguments")
	}
	if args[0].Typ != resp.TypeBulk || args[1].Typ != resp.TypeBulk {
		return resp.Value{}, fmt.Errorf("Invalid argument types. Must be Bulk Strings")
	}

	key := args[0].Bulk
	field := args[1].Bulk
	val := args[2].Bulk

	hashDB.mutex.Lock()
	_, exists := hashDB.entries[key]
	if !exists {
		hashDB.entries[key] = map[string]string{}
	}
	hashDB.entries[key][field] = val
	hashDB.mutex.Unlock()

	return resp.Value{Typ:resp.TypeString, Str: "OK"}, nil
}

func handleHGet(args []resp.Value) (resp.Value, error) {
	if len(args) != 2{
		return resp.Value{}, fmt.Errorf("Invalid Number of arguments")
	}
	if args[0].Typ != resp.TypeBulk {
		return resp.Value{}, fmt.Errorf("Invalid argument types. Must be Bulk Strings")
	}

	key := args[0].Bulk
	field := args[1].Bulk

	hashDB.mutex.RLock()
	val, exists := hashDB.entries[key][field]
	hashDB.mutex.RUnlock()
	
	if !exists {
		return resp.Value{Typ:resp.TypeNull}, nil
	}

	return resp.Value{Typ:resp.TypeBulk, Bulk: val}, nil
}
