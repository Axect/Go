package main

import (
	"fmt"

	. "github.com/Axect/Go/Numerical/Lagrange"
)

func main() {
	xseq := []float64{1., 2., 3., 4.}
	yseq := []float64{1., 4., 9., 16.}
	p := LPolynomial(xseq, yseq)
	fmt.Println(p(100))
}
