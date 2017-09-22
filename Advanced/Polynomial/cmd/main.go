package main

import (
	"fmt"

	"github.com/Axect/Go/Advanced/Polynomial"
)

func main() {
	var A, B Polynomial.Polynomial
	A.Coeff = []float64{1, -2, 3, 4}
	B.Coeff = []float64{2, 1}
	C := Polynomial.Add(A, B)
	D := Polynomial.Sub(A, B)
	C.Show()
	D.Show()
	E := Polynomial.Diff(D)
	E.Show()

	var P, Q Polynomial.Polynomial
	P.Coeff = []float64{1, 3, 3, 1}
	Q.Coeff = []float64{1, 1}
	R, r := Polynomial.Honor(P, Q)
	R.Show()
	fmt.Println(r)
	r1, r2 := Polynomial.RealQuad(R)
	fmt.Println(r1, r2)
}
