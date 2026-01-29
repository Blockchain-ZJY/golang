package p2p

// Peer is interface that represents the remote nodeeeran
type Peer interface {
}

// ransport is anything that handles the communication
// between the nodes in the network.
type Transport interface {
	ListenAndAccept() error
}
