package main

import (
	"fmt"

	"github.com/Axect/Go/Tutorial/method"
)

func main() {
	A := method.Person{
		Name:    "Axect",
		Account: 1,
	}
	A.Add()
	fmt.Println(A)
	A.Subtract()
	fmt.Println(A)
}
