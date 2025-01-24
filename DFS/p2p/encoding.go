package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct{}

func (doc GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (doc DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1028)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	fmt.Println(string(buf[:n]))

	return nil
}
