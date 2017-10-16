package main

import (
	"fmt"
	"time"

	"github.com/Axect/Go/Package/Time"
	"github.com/Axect/Go/Package/array"
)

type Vector = []float64

func Sum1(V Vector) float64 {
	start := time.Now()
	defer Time.TimeTrack(start, "Sum1")
	s := 0.
	for _, elem := range V {
		s += elem
	}
	return s
}

func Sum2(V Vector) float64 {
	start := time.Now()
	defer Time.TimeTrack(start, "Sum2")
	var L [1E+08]float64
	copy(L[:], V)
	s := 0.
	Q := &L
	for _, elem := range Q {
		s += elem
	}
	return s
}

func main() {
	V := array.Create(1, 1, 100000000)
	fmt.Println(Sum1(V))
	fmt.Println(Sum2(V))
}
