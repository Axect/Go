package root

import (
	"fmt"
	"math"
)

const h = 1e-09

// Normal is normal type of function for convenience
type Normal func(float64) float64

// Differentiate is differentiate
func Differentiate(f Normal, x float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}

// Derivative is derivative
func Derivative(f Normal) Normal {
	return func(x float64) float64 {
		return Differentiate(f, x)
	}
}

// Error check error
func Error(app, ans float64) float64 {
	return (app - ans)
}

// NewtonErr calculates number of iteration below given error by Newton's method
func NewtonErr(f Normal, x float64, ans float64) float64 {
	df := Derivative(f)
	var err float64
	i := 1
	for i < 100 {
		x -= f(x) / df(x)
		err = Error(x, ans)
		if math.Abs(err) == 0 {
			break
		}
		i++
	}
	fmt.Println("Number of Iterations : ", i)
	fmt.Println("Error is : ", err)
	return x
}

// Newton is newton's method
func Newton(f Normal, x float64, n int) float64 {
	df := Derivative(f)
	i := 1
	for i < n {
		x -= f(x) / df(x)
		i++
	}
	fmt.Println("Number of Iterations : ", i)
	return x
}
