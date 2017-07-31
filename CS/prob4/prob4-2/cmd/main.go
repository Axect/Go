package main

import (
	"fmt"
	"math"

	"github.com/Axect/Go/CS/prob4/prob4-2"
)

// TestFunc is just test function
func TestFunc(x float64) float64 {
	return math.Pow(x, 2) - 4
}

func TestFunc2(x float64)

func main() {
	fmt.Println("x^2 - 4 = 0")
	fmt.Println(root.NewtonErr(TestFunc, 1, 2))
	fmt.Println(root.Newton(TestFunc, 1, 10))
	fmt.Println("------------------------------")
}
