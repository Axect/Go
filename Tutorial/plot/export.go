package main

import (
	"fmt"
	"math"

	"github.com/Axect/Go/Package/csv"
)

// X is just test list
var X [100]float64

// Y is also
var Y [100]float64

// MakeSeq makes seq
func MakeSeq() (*[100]float64, *[100]float64) {
	Temp := &X
	Temp2 := &Y
	for i := range Temp {
		j := float64(i)
		Temp[i] = j / 10
	}
	for i, elem := range Temp {
		Temp2[i] = math.Sin(elem)
	}
	return &X, &Y
}

func main() {
	seq, seq2 := MakeSeq() // return Pointer of original array
	Y := make([][]string, 100, 100)
	for i, elem := range seq {
		temp := fmt.Sprint(elem)
		temp2 := fmt.Sprint(seq2[i])
		q := []string{temp, temp2}
		Y[i] = q
	}
	csv.Write(Y, "Data/test")
}
