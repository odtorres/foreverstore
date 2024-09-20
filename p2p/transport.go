package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
	Close() error
}

// transport is any thing that handles the communication
// between the nodes in the network. This can be of the form (TCM UDP, websocket, etc)
type Transport interface {
	listenAndAccept() error
	Consume() <-chan RPC
}
