package main

import (
	"fmt"
)

// =============================================================================
// Types
// =============================================================================

// ContraVector is Contravariant Vector
type ContraVector [4]float64

// CoVector is Covariant Vector
type CoVector ContraVector

// Tensor is Tensor
type Tensor [4][4]float64

// Metric Tensor
type Metric Tensor

// =============================================================================
// Methods
// =============================================================================

// Show is show
func (g Metric) Show() {
	fmt.Printf("diag[%v, %v, %v, %v]", g[0][0], g[1][1], g[2][2], g[3][3])
}

// =============================================================================
// Main Function
// =============================================================================
func main() {
	A := ContraVector{1, 0, 0, 0}
	g := Metric{{1, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, -1, 0}, {0, 0, 0, -1}}
	fmt.Println(A)
	g.Show()
}
