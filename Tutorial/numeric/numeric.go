package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/Axect/Numeric/Polynomial"
	"github.com/Axect/Numeric/array"
	"github.com/Axect/Numeric/spline"
	"github.com/Axect/Numeric/stats"
)

// Vector is type alias
type Vector = array.Vector

var wg sync.WaitGroup

func main() {
	// Create X Array
	X := array.Create(-5, 1, 5)
	// Create Polynomial
	var P Polynomial.Polynomial
	P.Coeff = Vector{1, 2, -3}
	// Make Function
	f := P.PolyFunc()
	// Create Y Array
	Y := make(Vector, len(X), len(X))
	for i := range Y {
		Y[i] = f(X[i])
	}
	// Spline Object
	spl := spline.NewCubic(X, Y)
	// Evaluate
	T := array.Create(-5, .1, 5)
	S := spl.Evaluate(T)

	// Calc Correlation Coeff
	c1 := stats.Cor(X, Y)
	c2 := stats.Cor(T, S)
	fmt.Println("Correlation: ", c1, c2)
	C1 := stats.CovMatrix(X, Y)
	C2 := stats.CovMatrix(T, S)
	array.MatrixForm(C1)
	array.MatrixForm(C2)

	array.Write([]Vector{T, S}, "Test.csv")

	wg.Add(1)
	go Routine()
	wg.Wait()
	fmt.Println()
	fmt.Println("All Process Finished!")
}

// Routine for Julia
func Routine() {
	defer wg.Done()

	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command("julia", "plot.jl").Output(); err != nil {
		panic(err)
	}
	str := string(cmdOut)
	fmt.Println(str)
	fmt.Println("Plot Finished!")
	return
}
