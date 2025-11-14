package main
import (
	"os"
	"github.com/Drime648/coding-challenges/redis/internal/resp"
	"sync"
	"time"
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



