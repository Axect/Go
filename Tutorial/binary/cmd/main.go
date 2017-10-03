package main

import (
	"fmt"
	"log"

	. "github.com/Axect/Go/Tutorial/binary"
)

func main() {
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value'", values[i], "': ", err)
		}
	}

	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Search(s)
	if !found {
		log.Fatal("Cannot Find '", s, "'")
	}
	fmt.Print("Found ", s, ": '", d, "'\n")
}
