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

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddress  string
	HandshakerFunc HandshakerFunc
	Decoder        Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// Consume returns a channel that can be used to receive RPC messages
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenerAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
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

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakerFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP Handshake error: %s\n", err)
		return
	}

	rpc := RPC{}
	for {
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP Error decoding message: %s\n", err)
			fmt.Printf("Closing connection from %+v\n", conn.RemoteAddr())
			return
		}
		rpc.From = conn.RemoteAddr()
		fmt.Printf("Received message: %+v\n", rpc)
	}
}
