package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/amanbhardwaj12072003/Distributed_File_Storage/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(),
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}
	s := NewFileServer(fileServerOpts)
	tcpTransport.OnPeer = s.OnPeer

	return s
}

// /*

func main() {
	server_1 := makeServer(":3000", "")
	server_2 := makeServer(":6000", "")
	server_3 := makeServer(":5000", ":3000", ":6000")

	go func() { log.Fatal(server_1.Start()) }()
	time.Sleep(2 * time.Second)
	go func() { log.Fatal(server_2.Start()) }()

	go server_3.Start()
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {

		key := fmt.Sprintf("cowpicture_%d.jpg", i)
		// key := "catspicture.jpg"
		data := bytes.NewReader([]byte("my big data file here!"))
		server_3.Store(key, data)

		if err := server_3.store.Delete(key); err != nil {
			log.Fatal(err)
		}

		r, err := server_3.Get(key)
		if err != nil {
			log.Fatal(err)
		}

		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}

}

// */

/*

func main() {
	server_1 := makeServer(":3000", "")
	server_2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(server_1.Start())
	}()
	time.Sleep(1 * time.Second)

	go server_2.Start()
	time.Sleep(1 * time.Second)

	key := "catspicture.jpg"
	data := bytes.NewReader([]byte("my big data file here!"))
	server_2.Store(key, data)

	if err := server_2.store.Delete(key); err != nil {
		log.Fatal(err)
	}

	r, err := server_2.Get(key)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

*/
