package p2p

// Peer is an interface that represents the remote node 
type Peer interface{
	Close() error
}

// Transport is anything that handles the communication 
// between the nodes in the network. Which can be of 
// form (TCP, UDP, Websockets, etc)
type Transport interface{
	ListenAndAccept() error
	Consume() <- chan RPC
}