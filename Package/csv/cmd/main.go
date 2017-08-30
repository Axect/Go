package main

import (
	"fmt"

	"github.com/Axect/Go/Package/csv"
)

func main() {
	Txt := [][]string{[]string{"1", "2", "3"}, []string{"4", "5", "6"}}
	csv.Write(Txt, "test.csv")
	A := csv.Read("test.csv")
	B := csv.Convert(A)
	fmt.Println(B)
}
