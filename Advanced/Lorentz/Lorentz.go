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

// Converter is CO -> Contra
func Converter(A CoVector) ContraVector {
	var B ContraVector
	B.t = A.t
	B.x = A.x
	B.y = A.y
	B.z = A.z
	return B
}

// Raising is Co -> Contra
func (B *CoVector) Raising() {
	B.x = -B.x
	B.y = -B.y
	B.z = -B.z
}

// Contraction is contraction
func Contraction(A ContraVector, B CoVector) float64 {
	g := []float64{1, -1, -1, -1}
	return g[0]*A.t*B.t + g[1]*A.x*B.x + g[2]*A.y*B.y + g[3]*A.z*B.z
}

func main() {
	A := ContraVector{
		t: 1,
		x: 1,
		y: 2,
		z: 3,
	}
	B := CoVector{
		t: 2,
		x: 0,
		y: 5,
		z: 6,
	}
	C := Contraction(A, B)
	fmt.Println(C)
	B.Raising()
	fmt.Println(B)
}
