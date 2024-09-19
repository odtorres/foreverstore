package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn

	//if we dial and retreive a conn => outbound is true
	//if we accept and retreive a conn => outbound is false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakerFunc
	decoder       Decoder

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenerAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept errir:%s\n", err)
		}

		fmt.Printf("Accepted connection from %+v\n", conn)

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {
		fmt.Printf("Handshake error: %s\n", err)
		return
	}

	lenDecadeError := 0
	//read loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			lenDecadeError++
			if lenDecadeError > 3 {
				fmt.Printf("Too many decode errors, closing connection: %s\n", err)
				return
			}
			fmt.Printf("TCP Error decoding message: %s\n", err)
			continue
		}
	}

	fmt.Printf("Handling connection from %+v\n", peer)
}
