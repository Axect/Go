package main

import (
	"fmt"

	"github.com/Axect/Go/Package/latexize"
)

func main() {
	A := "L = -1/2 * m^2 * \\phi^2"
	B := latexize.LaTeXize(A)
	fmt.Println(B)
}
