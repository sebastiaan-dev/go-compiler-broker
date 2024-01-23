//
//  Majordomo Protocol worker example.
//  Uses the mdwrk API to hide all MDP aspects
//

package main

import (
	"log"
	"os"

	generated "github.com/sebastiaan-dev/go-compiler-broker/pkg/serialization"

	"capnproto.org/go/capnp/v3"
)

func parseMessage(revMsg []byte) (option string, msgBody [][]byte) {
	msg, err := capnp.Unmarshal(revMsg)
	if err != nil {
		panic(err)
	}
	book, err := generated.ReadRootBook(msg)
	if err != nil {
		panic(err)
	}
	book.Title()

	return
}

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, _ := NewMdwrk("tcp://localhost:5555", "echo", verbose)

	var err error
	var request, reply [][]byte
	for {
		request, err = session.Recv(reply)
		if err != nil {
			break //  Worker was interrupted
		}

		reply = request //  Echo is complex... :-)
	}
	log.Println(err)
}
