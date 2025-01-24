package main

import (
	"dfs/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":3000") //Creates a new TCP trnsport with listen address "3000"
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err) //If any error occurs log it and terminate the program
	}

	select {} //Blocks the program from exiting
}
