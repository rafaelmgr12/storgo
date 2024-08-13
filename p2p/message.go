package p2p

// Message represents any arbitraty data that is being sent over the
// each transport bewteen two nodes in the network
type Message struct {
	Payload []byte
}
