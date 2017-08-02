package main

import (
	"math"

	"github.com/Axect/Go/CS/prob4/prob4-3"
)

func main() {
	maxima.CheckMax(maxima.GSS(TestFunc, -4, 2))
}

// TestFunc for test
func TestFunc(x float64) float64 {
	return math.Pow((x - 1), 2)
}
