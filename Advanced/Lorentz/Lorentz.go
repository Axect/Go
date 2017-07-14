package main

import (
	"fmt"
)

// ContraVector is Contravariant Vector
type ContraVector struct {
	t float64
	x float64
	y float64
	z float64
}

// CoVector is Covariant Vector
type CoVector ContraVector

// Metric is Metric Tensor
type Metric struct {
	diag [4]float64
}

// g is a metric tensor which for flat space.
var g = Metric{
	diag: [4]float64{1., -1., -1., -1.},
}

// Raising is Co -> Contra
func (B CoVector) Raising() ContraVector {
	var C ContraVector
	C.t, C.x, C.y, C.z = g.diag[0]*B.t, g.diag[1]*B.x, g.diag[2]*B.y, g.diag[3]*B.z
	return C
}

// Lowering is Contra -> Co
func (B ContraVector) Lowering() CoVector {
	var C CoVector
	C.t, C.x, C.y, C.z = g.diag[0]*B.t, g.diag[1]*B.x, g.diag[2]*B.y, g.diag[3]*B.z
	return C
}

// Contraction is contraction
func Contraction(A ContraVector, B CoVector) float64 {
	return A.t*B.t + A.x*B.x + A.y*B.y + A.z*B.z
}

func main() {
	var a, b, c, d float64
	var p, q, r, s float64
	fmt.Print("Please Input 4-Components of Contravariant Vector: ")
	fmt.Scan(&a, &b, &c, &d)
	A := ContraVector{
		t: a,
		x: b,
		y: c,
		z: d,
	}
	fmt.Print("Please Input 4-Components of Covariant Vector: ")
	fmt.Scan(&p, &q, &r, &s)
	B := CoVector{
		t: p,
		x: q,
		y: r,
		z: s,
	}
	S := Contraction(A, B)
	C := B.Raising()
	D := A.Lowering()
	E := Contraction(C, D)
	fmt.Println(S, E)
}
