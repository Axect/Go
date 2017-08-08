package main

import (
	"fmt"

	. "github.com/Axect/Go/Package/Numeric/Lagrange"
	"github.com/Axect/Go/Package/array"
)

func main() {
	xseq := []float64{1., 2., 3., 4.}
	yseq := []float64{1., 4., 9., 16.}
	p := LPolynomial(xseq, yseq)
	X := array.Arithmetic(0., 0.01, 100)
	fmt.Println(X)
}
