package main

import (
	"log"
	"time"

	"capnproto.org/go/capnp/v3"
	zmq "github.com/pebbe/zmq4"
	generated "github.com/sebastiaan-dev/go-compiler-broker/pkg/serialization"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:5555")

	for {
		// Wait for next request from client
		revMsg, _ := s.RecvBytes(0)
		msg, err := capnp.Unmarshal(revMsg)
		if err != nil {
			panic(err)
		}
		book, err := generated.ReadRootBook(msg)
		if err != nil {
			panic(err)
		}
		title, _ := book.Title()

		log.Printf("Received %s\n", title)

		// Do some 'work'
		time.Sleep(time.Second * 1)

		// Send reply back to client
		s.Send("Accepted", 0)
	}
}
