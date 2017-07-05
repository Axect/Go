package main

import (
	"fmt"

	"github.com/Axect/Go/Tutorial/varinit"
)

//S, T are just Const
const (
	S = 10
	T = 12
)

func main() {
	fmt.Println(varinit.Add(S, T))
}
