package main

import (
	"fmt"
	"log"

	"github.com/odtorres/foreverstore/p2p"
)

func OnPeer(p2p.Peer) error {
	fmt.Println("Doing some logic with the peer outside of TCPTransport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress:  ":3000",
		HandshakerFunc: p2p.NOPHandshakeFunc,
		Decoder:        p2p.DefaultDecoder{},
		OnPeer:         OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			rpc := <-tr.Consume()
			fmt.Printf("Received RPC: %+v\n", rpc)
		}
	}()

	if err := tr.ListenerAndAccept(); err != nil {
		log.Fatalf("Error listening and accepting: %s", err)
	}

	fmt.Println("Listening on", tr.ListenAddress)

	select {}
}
