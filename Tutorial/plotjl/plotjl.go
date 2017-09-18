package main

import (
	"fmt"
	"log"
	"math"

	"github.com/Axect/Go/Package/array"
	"github.com/Axect/csv"
)

type Vector = []float64

func main() {
	X := array.Create(0, 0.1, 10)
	Y := make(Vector, len(X), len(X))
	for i := range Y {
		Y[i] = math.Sin(X[i])
	}
	Z := Convert(X, Y)
	csv.Write(Z, "sin.csv")
}

// Convert just converts two Vector to [][]string
func Convert(V, W Vector) [][]string {
	if len(V) != len(W) {
		log.Fatal("Lengths are not matched!")
	}
	A := make([][]string, len(V), len(V))
	for i := range A {
		A[i] = []string{fmt.Sprint(V[i]), fmt.Sprint(W[i])}
	}
	return A
}
