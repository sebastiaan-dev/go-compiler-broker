package main

import (
	"fmt"

	"github.com/sebastiaan-dev/go-compiler-broker/serialization/generated"

	"capnproto.org/go/capnp/v3"
)

func main() {
	arena := capnp.SingleSegment(nil)

	_, seg, err := capnp.NewMessage(arena)
	if err != nil {
		panic(err)
	}

	book, err := generated.NewRootBook(seg)
	if err != nil {
		panic(err)
	}

	_ = book.SetTitle("War and Peace")

	// Then, we set the page count.
	book.SetAuthor("Leo Tolstoy")

	// Finally, we "get" these fields and print them.
	title, _ := book.Title()
	author, _ := book.Author()

	fmt.Printf("%s by %s", title, author)
}
