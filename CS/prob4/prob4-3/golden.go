package maxima

import (
	"fmt"
	"math"
)

// Normal is type of normal function
type Normal func(float64) float64

// GR generate golden ratio number
func GR() float64 {
	return (-1. + math.Sqrt(5)) / 2
}

// CheckMax checks max
func CheckMax(list []float64, key bool) {
	if key {
		fmt.Println("Point of local minima : ", list[len(list)-1])
	} else {
		fmt.Println("Can't find local minima in this range")
		fmt.Println("The Last value is :", list[len(list)-1])
	}
}

// GSS is golden section algorithm in given interval [a, b]
func GSS(f Normal, a, b float64) ([]float64, bool) {
	/* Welcome to GSS */
	c := GR()
	key := false
	list := make([]float64, 100, 100)
	x1 := c*a + (1-c)*b // segment
	x2 := (1-c)*a + c*b // segment
	x := (x1 + x2) / 2
	list[0] = x
	for i := 1; i < 100; i++ {
		fx1 := f(x1)
		fx2 := f(x2)
		if fx1 < fx2 {
			b = x2
			x2 = x1
			x1 = c*a + (1-c)*b
		} else if fx1 > fx2 {
			a = x1
			x1 = x2
			x2 = (1-c)*a + c*b
		} else {
			fmt.Println("Hmm")
		}
		x = (a + b) / 2
		list[i] = x
		if math.Abs(b-a) < 1e-15 {
			key = true
			fmt.Println("Number of Iterations : ", i)
			for j := i + 1; j < 100; j++ {
				list[j] = x
			}
			break
		}
	}
	return list, key
}
