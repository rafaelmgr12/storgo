package p2p

// RPC represents any arbitraty data that is being sent over the
// each transport bewteen two nodes in the network
type RPC struct {
	From    string
	Payload []byte
}
