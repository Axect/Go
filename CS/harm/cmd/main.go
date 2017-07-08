package main

import (
	"fmt"

	"github.com/Axect/Go/CS/harm"
)

func main() {
	Estr := harm.Harmonic()
	Ereal := harm.Converter(Estr)
	fmt.Println("Energy (symbolic): ", Estr)
	fmt.Println("Energy (numeric): ", Ereal)
}
