package main

import (
	"fmt"

	. "github.com/Axect/Go/Package/Numeric/Lagrange"
)

func main() {
	X := []float64{1., 2., 3., 4.}
	Y := []float64{1., 4., 9., 16.}
	L := LPolynomial(X, Y)
	fmt.Println(L(2.5))
}
