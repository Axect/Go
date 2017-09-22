package Polynomial

import (
	"log"
	"math"
)

// Honor represents Honor's Algorithm
func Honor(P, Q Polynomial) (Polynomial, float64) {
	if len(Q.Coeff) != 2 {
		log.Fatal("Just put linear polynomial")
	}
	nom := -Q.Coeff[0] / Q.Coeff[1]
	var R Polynomial
	var r float64
	R.Coeff = make([]float64, len(P.Coeff)-1, len(P.Coeff)-1)

	for i := range P.Coeff {
		switch {
		case i == 0:
			R.Coeff[i] = P.Coeff[i]
		case i == len(P.Coeff)-1:
			temp := nom*R.Coeff[i-1] + P.Coeff[i]
			r = temp
		default:
			temp := nom*R.Coeff[i-1] + P.Coeff[i]
			R.Coeff[i] = temp
		}
	}
	return R, r
}

// RealQuad is root formula
func RealQuad(P Polynomial) (float64, float64) {
	L := len(P.Coeff)
	if L != 3 {
		log.Fatal("Only quadratic formula supported")
	}
	a, b, c := P.Coeff[0], P.Coeff[1], P.Coeff[2]
	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		log.Fatal("There is no real root")
	}

	r1 := (-b + math.Sqrt(D)) / (2 * a)
	r2 := (-b - math.Sqrt(D)) / (2 * a)

	return r1, r2
}
