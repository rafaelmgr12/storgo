package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/rafaelmgr12/storgo/internal/cryptoutil"
	"github.com/rafaelmgr12/storgo/internal/server"
	"github.com/rafaelmgr12/storgo/internal/store"
	"github.com/rafaelmgr12/storgo/pkg/p2p"
)

func makeServer(listenAddr string, nodes ...string) *server.FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := server.FileServerOpts{
		EncKey:            cryptoutil.NewEncryptionKey(),
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: store.CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := server.NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}
func main() {

	s1 := makeServer(":3000", "")
	s2 := makeServer(":7000", "")
	s3 := makeServer(":5000", ":3000", ":7000")

	go func() {
		log.Fatal(s1.Start())
	}()

	time.Sleep(500 * time.Millisecond)
	go func() {
		log.Fatal(s2.Start())
	}()

	time.Sleep(2 * time.Second)

	go s3.Start()

	key := "picture.png"
	data := bytes.NewReader([]byte("my big data file here!"))
	s3.Store(key, data)

	r, err := s3.Get(key)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	time.Sleep(2 * time.Second)

	if err := s3.Delete(key); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)

}
