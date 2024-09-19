package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is a test function for the TCPTransport struct
func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)

	assert.Nil(t, tr.ListenerAndAccept())
}
