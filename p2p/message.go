package p2p

import "net"

// Message is a struct that represents the message that is sent between the nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}
