package p2p

//Message holds any arbitary data that is bieng sent over
//each transport between two nodes in the network
type Message struct {
	Payload []byte
}
