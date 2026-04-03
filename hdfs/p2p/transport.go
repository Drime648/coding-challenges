package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
}

// Transport is anything that handles communication between networked nodes.
// This can be of the form of TCP, UDP, Websockets, etc.
type Transport interface {
	ListenAndAccept() error
}
