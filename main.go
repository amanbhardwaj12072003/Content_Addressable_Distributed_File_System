package main

import (
	"bytes"
	"log"
	"time"

	"github.com/amanbhardwaj12072003/Distributed_File_Storage/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// Todo: onpeer func
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}
	s := NewFileServer(fileServerOpts)
	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	server_1 := makeServer(":3000", "")
	server_2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(server_1.Start())
	}()
	time.Sleep(1 * time.Second)

	go server_2.Start()
	time.Sleep(1 * time.Second)

	data := bytes.NewReader([]byte("my big data file here!"))

	server_2.StoreData("myprivatekey", data)

	select {}
}
