package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer reoresents the remote node over a TCP established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn

	//if we dial and retrieve conn => outbound == ture
	//if we accept and retrieve conn => outbound == false
	outbound bool
}

// This constructor creates new TCPPeer instance with the connection and weather it's inbound or outbound
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string       //The address and the port the transport will listen to
	listener      net.Listener // This is the network listener that waits for incoming TCP connections

	mu    sync.RWMutex
	peers map[net.Addr]Peer //A map to store peers identified by thier network address
}

// This constructor initializes a new TCPTransport instance with the listening address
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

// This method starts the TCP listener on the specified address t.listenaddress
func (t *TCPTransport) ListenAndAccept() error {
	var err error

	//net.Listen("tcp", t.listenAddress) called to bind the address and listen for incoming connections
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	//Spins of a go routine to accecpt connections councurrently
	go t.startAcceptLoop()

	return nil
}

// This method is running a seperate go routine to accept incoming connections continiously
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept() //Called to wait for incomming connections
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err) //if error occurs prints the following message
		}
		go t.handleConn(conn) //Each connection is handled in a new goroutine
	}
}

// This method is called when a new connection is accepted
// Creates a new TCPPeer instance to represent the peer and the logs the new connection
func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incomming connection %+v\n", peer)
}
