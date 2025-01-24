package p2p

//Peer is an interface that represents remote node.
type Peer interface{}

//Transport is anything that handles communication between nodes in the network.
//This can be of any type like TCP, UDP, websocket....
type Transport interface {
	ListenAndAccept() error //This starts the listening process for incomming connections.
}
