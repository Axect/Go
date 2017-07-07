package timetest

import (
	"fmt"
	"time"
)

// IntFunc is type of general functions
type IntFunc func(int) int

// FloatFunc is type of general functions
type FloatFunc func(float64) float64

// TestInt is test for int function
func TestInt(f IntFunc, x int) {
	start := time.Now()
	f(x)
	elapsed := time.Since(start)
	fmt.Println("Time Elapsed:", elapsed)
}

// TestFloat is test for float function
func TestFloat(f FloatFunc, x float64) {
	start := time.Now()
	f(x)
	elapsed := time.Since(start)
	fmt.Println("Time Elapsed:", elapsed)
}
