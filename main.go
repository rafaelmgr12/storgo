package main

import (
	"log"

	"github.com/rafaelmgr12/storgo/p2p"
)

func main() {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// TODO: onPeer func
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		ListenAddr:        ":3000",
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTansformFunc,
		Transport:         tcpTransport,
	}

	fs := NewFileServer(fileServerOpts)

	if err := fs.Start(); err != nil {
		log.Fatal(err)
	}

	select {}

}
