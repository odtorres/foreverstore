package main

import (
	"log"
	"time"

	"github.com/odtorres/foreverstore/p2p"
)

func main() {

	tcpTransportOPts := p2p.TCPTransportOpts{
		ListenAddress:  ":3000",
		HandshakerFunc: p2p.NOPHandshakeFunc,
		Decoder:        p2p.DefaultDecoder{},
		//todo on peer func
	}

	tcpTransport := p2p.NewTCPTransport(tcpTransportOPts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}

	s := NewFileServer(fileServerOpts)

	go func() {
		time.Sleep(3 * time.Second)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
