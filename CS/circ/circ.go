package main

import (
	"fmt"
	"math"

	"github.com/Axect/Go/Package/csv"
)

// MakeSeq makes Seq
func MakeSeq() (*[10001]float64, *[10001]float64, *[10001]float64) {
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

func main() {
	A, B, C := MakeSeq()
	Seq := make([][]string, len(A), len(A))
	for i := range A {
		q1, q2, q3 := fmt.Sprint(A[i]), fmt.Sprint(B[i]), fmt.Sprint(C[i])
		str := []string{q1, q2, q3}
		Seq[i] = str
	}
	csv.Write(Seq, "Data/Circ.csv")
}
