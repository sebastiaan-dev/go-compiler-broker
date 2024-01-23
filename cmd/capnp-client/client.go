package main

import (
	"fmt"

	generated "github.com/sebastiaan-dev/go-compiler-broker/pkg/serialization"

	"capnproto.org/go/capnp/v3"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	s.Connect("tcp://localhost:5555")

	// Do 10 requests, waiting each time for a response
	for i := 0; i < 10; i++ {
		fmt.Printf("Sending request %d...\n", i)

		arena := capnp.SingleSegment(nil)
		msg, seg, err := capnp.NewMessage(arena)
		if err != nil {
			panic(err)
		}

		book, err := generated.NewRootBook(seg)
		if err != nil {
			panic(err)
		}

		_ = book.SetTitle("War and Peace")

		book.SetAuthor("Leo Tolstoy")

		b, err := msg.Marshal()
		if err != nil {
			panic(err)
		}

		s.SendBytes(b, 0)

		revMsg, _ := s.Recv(0)
		fmt.Printf("Received reply %d [ %s ]\n", i, revMsg)
	}
}
