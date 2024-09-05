package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface {
	net.Conn
	Send([]byte) error
}

// Tansport is anything that handlers the communication
// between the nodes in the network. This can be of the
// form (TCP,UDP websockets, ....)
type Transport interface {
	Addr() string
	ListenAndAccept() error
	Dial(string) error
	Consume() <-chan RPC
	Close() error
}
