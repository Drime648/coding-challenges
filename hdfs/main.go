package main

import (
	"fmt"
	"time"

	"github.com/Drime648/coding-challenges/hdfs/p2p"
)

func onPeer(p2p.Peer) error {
	fmt.Println("Running Peer logic")
	return nil
}

func main() {

	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NopHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// OnPeer: onPeer,
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_files",
		PathTransformFunc: CASPathTransformFun,
		Transport:         tcpTransport,
	}
	s := NewFileServer(fileServerOpts)

	go func() {
		time.Sleep(3 * time.Second)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		panic(err)
	}




}
