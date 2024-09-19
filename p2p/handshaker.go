package p2p

// HandshakerFunc is a function that is used to handshake with the remote node
type HandshakerFunc func(Peer) error

// NOPHandshakeFunc is a no-op handshake function that always returns nil
func NOPHandshakeFunc(Peer) error {
	return nil
}
