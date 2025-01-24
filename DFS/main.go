package main

import (
	"dfs/p2p"
	"log"
)

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandShakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tr := p2p.NewTCPTransport(tcpOpts) //Creates a new TCP trnsport with listen address "3000"
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err) //If any error occurs log it and terminate the program
	}

	select {} //Blocks the program from exiting
}
