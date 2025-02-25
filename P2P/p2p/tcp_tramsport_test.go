package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is a unit test for TCPTransport
func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)

	assert.Equal(t, tr.listenAddress, listenAddr) //Ckeck if the listen address matches

	assert.Nil(t, tr.ListenAndAccept()) //Verify if ListenAndAccept returns no error

	select {} //Used to block indefinity, keeping the program running
}
