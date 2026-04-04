package p2p

import "net"

// RPC holds any arbitrary data that is
// sent over each transport between 2 nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}
