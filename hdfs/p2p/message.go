package p2p


// Message holds any arbitrary data that is
// sent over each transport between 2 nodes
type Message struct {
	Payload []byte
}
