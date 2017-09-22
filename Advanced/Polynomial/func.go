package Polynomial

import "math"

// PF for convenience
type PF func(float64) float64

// PolyFunc makes function
func (P Polynomial) PolyFunc() PF {
	L := len(P.Coeff) - 1
	return func(x float64) float64 {
		s := 0.
		for i := range P.Coeff {
			s += P.Coeff[i] * math.Pow(x, float64(L-i))
		}
		return s
	}
}
