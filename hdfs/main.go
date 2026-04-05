package main

import (
	"log"

	"github.com/Drime648/coding-challenges/hdfs/p2p"
)

// func onPeer(p2p.Peer) error {
// 	fmt.Println("Running Peer logic")
// 	return nil
// }

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NopHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// OnPeer: onPeer,
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFun,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}
	return NewFileServer(fileServerOpts)
}

func main() {

	s1 := makeServer(":3000")
	s2 := makeServer(":4000", ":3000")
	go func() {
		log.Fatal(s1.Start())
	}()

	s2.Start()

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	s.Stop()
	// }()


}
