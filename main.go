package main

import (
	"log"

	"github.com/chrollo-lucifer-12/file-storage/p2p"
)

func main() {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_objects",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}

	s := NewFileServer(fileServerOpts)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
