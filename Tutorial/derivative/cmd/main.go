package main

import (
	"fmt"
	"math"

	"github.com/Axect/Go/Tutorial/derivative"
)

// y is just test function
func y(x float64) float64 {
	return math.Pow(x, 3)
}

func main() {
	g := derivative.Derivative(y)
	fmt.Println(g(3))
}
