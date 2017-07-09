package main

import (
	"fmt"

	"github.com/Axect/Go/Package/csv"
)

// X is just test list
var X [100]float64

// MakeSeq makes seq
func MakeSeq() *[100]float64 {
	Temp := &X
	for i := range Temp {
		Temp[i] = float64(i + 1)
	}
	return &X
}

func main() {
	seq := MakeSeq()
	Y := make([][]string, 100, 100)
	for i, elem := range seq {
		temp := fmt.Sprint(elem)
		q := []string{temp}
		Y[i] = q
	}
	csv.Write(Y, "test")
}
