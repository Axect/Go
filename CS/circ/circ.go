package main

import (
	"fmt"
	"math"
	"os"

	"github.com/Axect/Go/Package/csv"
)

// =============================================================================
// Make Sequence
// =============================================================================

// MakeSeqCirc makes Seq
func MakeSeqCirc() (*[10001]float64, *[10001]float64, *[10001]float64) {
	var X, YUp, YDown [10001]float64
	Temp1, Temp2, Temp3 := &X, &YUp, &YDown
	for i := range Temp1 {
		j := float64(i-5000) / 5000
		Temp1[i] = j
		Temp2[i] = math.Sqrt(1 - math.Pow(j, 2))
		Temp3[i] = -math.Sqrt(1 - math.Pow(j, 2))
	}
	return &X, &YUp, &YDown
}

// MakeSeqTangent makes Seq, xP + yQ = 1
func MakeSeqTangent(x, y float64) (*[10001]float64, *[10001]float64) {
	if y == 0 {
		fmt.Println("Error: y can't be zero")
		os.Exit(1)
	}
	var P, Q [10001]float64
	Temp1, Temp2 := &P, &Q
	for i := range Temp1 {
		j := float64(i-5000) / 2000
		Temp1[i] = j
		Temp2[i] = -x/y*j + 1/y
	}
	return &P, &Q
}

// =============================================================================
// Main Function
// =============================================================================

func main() {
	A, B, C := MakeSeqCirc()
	Seq := make([][]string, len(A), len(A))
	for i := range A {
		q1, q2, q3 := fmt.Sprint(A[i]), fmt.Sprint(B[i]), fmt.Sprint(C[i])
		str := []string{q1, q2, q3}
		Seq[i] = str
	}
	csv.Write(Seq, "Data/Circ.csv")

	P, Q := MakeSeqTangent(math.Sqrt(2)/2, math.Sqrt(2)/2)
	Seq2 := make([][]string, len(P), len(P))
	for i := range P {
		r1, r2 := fmt.Sprint(P[i]), fmt.Sprint(Q[i])
		str2 := []string{r1, r2}
		Seq2[i] = str2
	}
	csv.Write(Seq2, "Data/Tan.csv")
}
