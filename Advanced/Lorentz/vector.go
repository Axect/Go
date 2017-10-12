package Lorentz

import (
	"fmt"
	"math"
)

// VectorType is interface for vector
type VectorType = string

// Vector is mother type for all vectors
type Vector struct {
	Type VectorType
	Comp []float64
}

// Norm is method for Vector
func (V *Vector) Norm() float64 {
	s := 0.
	for _, elem := range V.Comp {
		s += math.Pow(elem, 2.)
	}
	return math.Sqrt(s)
}

// Show shows vector type and components
func (V *Vector) Show() {
	fmt.Printf("%d-%s\n", len(V.Comp), V.Type)
	fmt.Println(V.Comp)
}
