package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	WSPort = ":3223"
)

type MsgType string

const (
	MsgType_Broadcast MsgType = "broadcast"
)

type ReqMsg struct {
	MsgType MsgType
	Client  *Client
	Data    string
}
type RespMsg struct {
	MsgType  MsgType
	Data     string
	SenderID string
}

func NewRespMsg(msg *ReqMsg) *RespMsg {
	return &RespMsg{
		MsgType:  msg.MsgType,
		Data:     msg.Data,
		SenderID: msg.Client.ID,
	}
}

type Client struct {
	ID string

	mu   *sync.RWMutex
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	ID := rand.Text()[:9]
	return &Client{
		ID:   ID,
		mu:   new(sync.RWMutex),
		conn: conn,
	}
}

func (c *Client) readMsgLoop(srv *Server) {
	defer func() {
		c.conn.Close()
		srv.leaveServerCH <- c
	}()
	for {
		_, b, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		msg := new(ReqMsg)
		if err = json.Unmarshal(b, msg); err != nil {
			fmt.Printf("unable to unmarshal msg %v\n", err)
			continue
		}

		srv.broadcastCH <- msg
	}
}

type Server struct {
	mu            *sync.RWMutex
	clients       map[string]*Client
	joinServerCH  chan *Client
	leaveServerCH chan *Client
	broadcastCH   chan *ReqMsg
}

func NewServer() *Server {
	return &Server{
		mu:            new(sync.RWMutex),
		clients:       map[string]*Client{},
		joinServerCH:  make(chan *Client, 64),
		leaveServerCH: make(chan *Client, 64),
		broadcastCH:   make(chan *ReqMsg, 64),
	}
}

func (s *Server) HandleWS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  512,
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
	s.joinServerCH <- client

	go client.readMsgLoop(s)
}

func (s *Server) AcceptLoop() {
	for {
		select {
		case c := <-s.joinServerCH:
			s.joinServer(c)
		case c := <-s.leaveServerCH:
			s.leaveServer(c)
		case msg := <-s.broadcastCH:
			go s.broadcast(msg)
		}
	}
}

func (s *Server) joinServer(c *Client) {
	s.clients[c.ID] = c
	fmt.Printf("client joined server, cID = %s\n", c.ID)
}

func (s *Server) leaveServer(c *Client) {
	delete(s.clients, c.ID)
	fmt.Printf("client left server, cID = %s\n", c.ID)
}

func (s *Server) broadcast(msg *ReqMsg) {
	cls := []*Client{}
	//Want to make a snapshot of the clients before sending
	s.mu.RLock()
	for _, c := range s.clients {
		if c.ID != msg.Client.ID {
			cls = append(cls, c)
		}
	}
	s.mu.RUnlock()
	resp := NewRespMsg(msg)
	for _, c := range cls {
		if err := c.conn.WriteJSON(resp); err != nil {
			fmt.Printf("Error sending message to client ID %s. err = %v\n", c.ID, err)
			continue
		}
	}

	fmt.Println("Broadcast has been sent.")
}

func CreateWSServer() {
	s := NewServer()
	go s.AcceptLoop()
	http.HandleFunc("/", s.HandleWS)

	fmt.Printf("Starting server on port: %s\n", WSPort)

	log.Fatal(http.ListenAndServe(WSPort, nil))
}

//TODO
// [x] HTTP server
// [x] Upgrade it to WS once client connects
// [x] Add WS client
// [x] Add newly connected ws to server
// [x] Remove client on disconnect
// [] broadcast messages to all clients. No race conditions.

func main() {
	// CreateWSServer()
}
