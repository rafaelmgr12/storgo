package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/rafaelmgr12/storgo/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(),
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTansformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}
func main() {

	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())

	}()
	time.Sleep(2 * time.Second)

	go s2.Start()

	time.Sleep(2 * time.Second)

	key := "collPicture.jpg"

	data := bytes.NewReader([]byte("my cool picture is here"))
	s2.Store(key, data)
	time.Sleep(500 * time.Millisecond)

	time.Sleep(1 * time.Second)

	if err := s2.store.Delete(key); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	r, err := s2.Get(key)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

}
