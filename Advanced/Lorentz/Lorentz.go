package main

import (
	"fmt"
)

// =============================================================================
// Types
// =============================================================================

// ContraVector is Contravariant Vector
type ContraVector [4]Scalar

// CoVector is Covariant Vector
type CoVector ContraVector

// Tensor is Tensor
type Tensor [4][4]Scalar

// Metric Tensor
type Metric Tensor

// Scalar is float64
type Scalar float64

// =============================================================================
// Methods
// =============================================================================

// Show is show
func (g Metric) Show() {
	fmt.Printf("diag[%v, %v, %v, %v]\n", g[0][0], g[1][1], g[2][2], g[3][3])
}

// Contraction is contraction
func (g Metric) Contraction(C1 ContraVector, C2 CoVector) Scalar {
	S := g[0][0]*C1[0]*C2[0] + g[1][1]*C1[1]*C2[1] + g[2][2]*C1[2]*C2[2] + g[3][3]*C1[3]*C2[3]
	return S
}

// =============================================================================
// Main Function
// =============================================================================
func main() {
	A := ContraVector{1, 1, 2, 3}
	B := CoVector{1, 1, 2, 3}
	g := Metric{{1, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, -1, 0}, {0, 0, 0, -1}}
	g.Show()
	S := g.Contraction(A, B)
	fmt.Println(S)
}
