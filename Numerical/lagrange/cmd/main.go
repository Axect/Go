package main

import (
	"fmt"

	. "github.com/Axect/Go/Package/Numeric/Lagrange"
	"github.com/Axect/Go/Package/array"
	"github.com/Axect/Go/Package/csv"
)

func main() {
	X := []float64{1., 2., 3., 4.}
	Y := []float64{1., 4., 9., 16.}
	L := LPolynomial(X, Y)
	T := array.Arithmetic(0., 0.01, 400)
	Q := make([]float64, len(T), len(T))
	temp := make([][]string, len(T), len(T))
	for i, t := range T {
		Q[i] = L(t)
		temp[i] = []string{fmt.Sprint(t), fmt.Sprint(Q[i])}
	}
	csv.Write(temp, "Data/lag.csv")
}
