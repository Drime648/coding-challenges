package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
	Close() error
}

// Transport is anything that handles communication between networked nodes.
// This can be of the form of TCP, UDP, Websockets, etc.
type Transport interface {
	ListenAndAccept() error
	Consume() <- chan RPC
	Close() error
}
