package p2p

const (
	IncomingMessage = 0x1
	IncomingStream  = 0x2
)

// RPC represents any arbitraty data that is being sent over the
// each transport bewteen two nodes in the network
type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
