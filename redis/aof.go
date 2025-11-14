package main
import (
	"os"
	"github.com/Drime648/coding-challenges/redis/internal/resp"
	"sync"
	"time"
	"io"
	"strings"
	"fmt"
)

type Aof struct {
	file *os.File
	mutex sync.Mutex
}

func NewAof(path string) (*Aof, error) {
	f, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	aof := Aof{file: f,}

	go func () {

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			aof.mutex.Lock()
			aof.file.Sync()
			aof.mutex.Unlock()
		}

	}()

	return &aof, nil
}


func (aof *Aof) Close() error {
	aof.mutex.Lock()
	defer aof.mutex.Unlock()
	return aof.file.Close()
} 


func (aof *Aof) Write(value resp.Value) error {
	aof.mutex.Lock()
	defer aof.mutex.Unlock()

	bytes := value.Marshal()
	_, err := aof.file.Write(bytes)
	if err != nil {
		return err
	}
	return nil

}

func (aof *Aof) Read() error {

	aof.mutex.Lock()
	defer aof.mutex.Unlock()

	fmt.Println("Started Reading Aof...")

	respParser := resp.NewResp(aof.file)
	for {
		value, err := respParser.Read()
		if err != nil {
			if err == io.EOF{ //end loop because we are at end of file
				break
			}
			return err
		}
		//assume that the file has no errors in respect to command
		//because we only wrote when it was a successful command
		
		command := strings.ToUpper(value.Array[0].Bulk)
		callback, _ := Handlers[command]
		callback(value.Array[1:])
	}
	fmt.Println("Finished Reading Aof")

	return nil
}


