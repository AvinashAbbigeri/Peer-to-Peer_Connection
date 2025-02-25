package p2p

//Message holds any arbitary data that is bieng sent over
//each transport between two nodes in the network
//Represents the data exchanged between the peers
type Message struct {
	Payload []byte
}
