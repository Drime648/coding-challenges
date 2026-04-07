package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	WSPort = ":3223"
)


type Client struct {
	ID string

	mu *sync.RWMutex
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	ID := rand.Text()[:9]
	return &Client{
		ID: ID,
		mu: new(sync.RWMutex),
		conn: conn,
	}
}

type Server struct {
	mu      *sync.RWMutex
	clients []*Client
}

func NewServer() *Server {
	return &Server{
		mu: new(sync.RWMutex),
		clients: []*Client{},
	}
}

func (s *Server) HandleWS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader {
		ReadBufferSize: 512,
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error on HTTP connection upgrade: %v\n", err)
		return
	}

	client := NewClient(conn)

	s.mu.Lock()
	s.clients = append(s.clients, client)
	s.mu.Unlock()
}

func CreateWSServer() {
	s := NewServer()
	http.HandleFunc("/", s.HandleWS)

	fmt.Printf("Starting server on port: %s\n", WSPort)

	log.Fatal(http.ListenAndServe(WSPort, nil))
}

//TODO
// [x] HTTP server
// [x] Upgrade it to WS once client connects
// [] Add WS client
// [] Add newly connected ws to server
// [] Remove client on disconnect
// [] broadcast messages to all clients. No race conditions.

func main() {
	// CreateWSServer()
}
