package Polynomial

import (
	"fmt"
	"math"
)

// =============================================================================
// Type
// =============================================================================

// Polynomial is Polynomial
type Polynomial struct {
	Coeff []float64
}

// Vector is type alias
type Vector = []float64

// =============================================================================
// Method
// =============================================================================

// Show is show polynomial
func (P Polynomial) Show() {
	var str string
	l := len(P.Coeff) - 1
	for i := range P.Coeff {
		switch {
		case i == 0:
			temp := fmt.Sprintf("%vx^%d", P.Coeff[i], l-i)
			str += temp
		case i == l:
			if P.Coeff[i] > 0 {
				temp := fmt.Sprintf(" + %v", P.Coeff[i])
				str += temp
			} else if P.Coeff[i] < 0 {
				temp := fmt.Sprintf(" - %v", math.Abs(P.Coeff[i]))
				str += temp
			}
		case i == l-1:
			if P.Coeff[i] > 0 {
				temp := fmt.Sprintf(" + %vx", P.Coeff[i])
				str += temp
			} else if P.Coeff[i] < 0 {
				temp := fmt.Sprintf(" - %vx", math.Abs(P.Coeff[i]))
				str += temp
			}
		default:
			if P.Coeff[i] > 0 {
				temp := fmt.Sprintf(" + %vx^%d", P.Coeff[i], l-i)
				str += temp
			} else if P.Coeff[i] < 0 {
				temp := fmt.Sprintf(" - %vx^%d", math.Abs(P.Coeff[i]), l-i)
				str += temp
			}
		}
	}
	fmt.Println(str)
}

// =============================================================================
// Function
// =============================================================================

// Minus makes polynomial to -polynomial
func Minus(P Polynomial) Polynomial {
	L := len(P.Coeff)
	var Q Polynomial
	Q.Coeff = make([]float64, L, L)
	for i := range Q.Coeff {
		Q.Coeff[i] = -P.Coeff[i]
	}
	return Q
}
