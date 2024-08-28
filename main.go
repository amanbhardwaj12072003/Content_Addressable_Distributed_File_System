package main

import (
	"log"

	"github.com/amanbhardwaj12072003/Distributed_File_Storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
