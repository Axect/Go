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

func TestFunc2(x float64) float64 {
	return math.Pow(x, 5) - 2*math.Pow(x, 3) - 7*math.Pow(x, 2) + 14
}

func TestFunc3(x float64) float64 {
	return x - math.Cos(x)
}

func main() {
	fmt.Println("TEST")
	fmt.Println("------------------------------")
	fmt.Println("x^2 - 4 = 0")
	fmt.Println("------------------------------")
	root.CheckWell(root.Newton(TestFunc, 1, 10))
	fmt.Println("------------------------------")
	fmt.Println("x^5 - 7x^3 -2x^2 + 14 = 0")
	fmt.Println("------------------------------")
	root.CheckWell(root.Newton(TestFunc2, 1, 10000))
	fmt.Println("------------------------------")
	fmt.Println("x - cosx = 0")
	fmt.Println("------------------------------")
	root.CheckWell(root.Newton(TestFunc3, 1, 10))
	fmt.Println()
}
