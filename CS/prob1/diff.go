package prob1

import (
	"fmt"
	"math"
)

const h = 1e-06

// Normal type is for convenience
type Normal func(float64) float64

// Diff differentiate function
func Diff(f Normal, x float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}

// =============================================================================
// Define Test Functions
// =============================================================================

// f is x^2
func f(x float64) float64 {
	return math.Pow(x, 2)
}

// f2 is sin(x)
func f2(x float64) float64 {
	return math.Sin(x)
}

// f3 is e^x
func f3(x float64) float64 {
	return math.Exp(x)
}

// f4 is ln(x)
func f4(x float64) float64 {
	return math.Log(x)
}

// =============================================================================
// Do!
// =============================================================================

// DoDiff is DoDiff
func DoDiff() {
	fmt.Println("---------------------------------------------")
	fmt.Println("Differentiation")
	fmt.Println("---------------------------------------------")
	result := Diff(f, 0)
	result2 := Diff(f2, 0)
	result3, result4 := Diff(f3, 0), Diff(f4, 0)
	fmt.Println(result, result2, result3, result4)
	fmt.Println()
}
