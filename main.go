package main

import (
	"log"

	"github.com/odtorres/foreverstore/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenerAndAccept(); err != nil {
		log.Fatalf("Error listening and accepting: %s", err)
	}

	select {}
}
