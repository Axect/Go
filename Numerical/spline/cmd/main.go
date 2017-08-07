package main

import (
	"fmt"

	"github.com/Axect/Go/Numerical/spline"
	"github.com/Axect/Go/Package/csv"
)

func main() {
	X := []float64{1., 2., 3., 4., 5.}
	Y := []float64{1., 4., 9., 16., 25.}
	Sp := spline.NewCubic(X, Y)

	T := make([]float64, 40, 40)
	for i := 0; i < len(T); i++ {
		T[i] = 1. + 0.1*float64(i)
	}
	S := Sp.Evaluate(T)
	Temp := make([][]string, len(T), len(T))
	for i := range T {
		Temp[i] = []string{fmt.Sprint(T[i]), fmt.Sprint(S[i])}
	}
	csv.Write(Temp, "Data/Cubic.csv")
}
