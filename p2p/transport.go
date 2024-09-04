package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface {
	net.Conn
	Send([]byte) error
}

// Transport is anything that handles the communication
// between the nodes in the network. Which can be of
// form (TCP, UDP, Websockets, etc)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
	Dial(string) error
}
