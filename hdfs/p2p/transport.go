package p2p

import "net"

// Peer is an interface that represents the remote node.
type Peer interface {
	RemoteAddr() net.Addr
	Close() error
}

// Transport is anything that handles communication between networked nodes.
// This can be of the form of TCP, UDP, Websockets, etc.
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <- chan RPC
	Close() error
}
