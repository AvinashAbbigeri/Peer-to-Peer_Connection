package p2p

import (
	"encoding/gob"
	"io"
)

// Interface for decoding messages
type Decoder interface {
	Decode(io.Reader, *Message) error
}

// Decodes a message serialized with encoding/gob
type GOBDecoder struct{}

func (doc GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

// Reads raw bytes and stores them in the message's Payload field
type DefaultDecoder struct{}

func (doc DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1028)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:n]

	return nil
}
