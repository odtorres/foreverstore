package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is a test function for the TCPTransport struct
func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddress:  ":3000",
		HandshakerFunc: NOPHandshakeFunc,
		Decoder:        DefaultDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tr.ListenAddress, ":3000")

	assert.Nil(t, tr.ListenAndAccept())
}
