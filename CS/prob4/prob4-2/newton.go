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

// Newton is newton's method which generates slice and whether it can find root.
func Newton(f Normal, x float64, n int) ([]float64, bool) {
	df := Derivative(f)
	i := 1
	list := make([]float64, n, n)
	key := false
	list[0] = x
	var temp, err float64
	for i < n {
		temp = x
		x -= f(x) / df(x)
		err = Error(x, temp)
		list[i] = x
		if math.Abs(err) == 0 {
			key = true
			for j := i + 1; j < n; j++ {
				list[j] = x
			}
			break
		}
		i++
	}
	fmt.Println("Number of Iterations : ", i)
	return list, key
}

// CheckWell checks whether root is well defined
func CheckWell(list []float64, key bool) {
	if key {
		fmt.Println("root : ", list[len(list)-1])
	} else {
		fmt.Printf("Can't find root in range %v\n", len(list))
		fmt.Println("The Last value is : ", list[len(list)-1])
	}
}
