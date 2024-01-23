//
//  Majordomo Protocol client example - asynchronous.
//  Uses the mdcli API to hide all MDP aspects
//

package main

import (
	"fmt"
	"log"
	"os"

	generated "github.com/sebastiaan-dev/go-compiler-broker/pkg/serialization"

	"capnproto.org/go/capnp/v3"
)

func buildMessage() ([]byte, error) {
	arena := capnp.SingleSegment(nil)
	msg, seg, err := capnp.NewMessage(arena)
	if err != nil {
		return nil, err
	}

	book, err := generated.NewRootBook(seg)
	if err != nil {
		return nil, err
	}

	_ = book.SetTitle("War and Peace")

	book.SetAuthor("Leo Tolstoy")

	return msg.Marshal()
}

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, _ := NewMdcli2("tcp://localhost:5555", verbose)

	var count int
	for count = 0; count < 100; count++ {
		err := session.Send("echo", []byte{})
		if err != nil {
			log.Println("Send:", err)
			break
		}
	}
	for count = 0; count < 100; count++ {
		_, err := session.Recv()
		if err != nil {
			log.Println("Recv:", err)
			break
		}
	}
	fmt.Printf("%d replies received\n", count)
}
