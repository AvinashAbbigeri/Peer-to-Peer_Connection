package p2p

//HandShakeFunc is a type alias for a function that takes a Peer and returns an error.
//Allows you to define custom handshake logic.
type HandShakeFunc func(Peer) error

func NOPHandShakeFunc(Peer) error { return nil }
