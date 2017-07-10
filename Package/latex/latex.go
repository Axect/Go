package main

import (
	"fmt"
	"strings"
)

// Powize makes TeX code to Go Code
func Powize(n string) string {
	Tempt := ""
	if strings.Contains(n, "^") {
		Temp := strings.Split(n, "")
		index := strings.Index(n, "^")
		var Front, Back string
		for i := range Temp {
			if i < index {
				Front += Temp[i]
			} else if i > index {
				Back += Temp[i]
			}
		}
		Tempt += fmt.Sprintf("Pow(%v,%v)", Front, Back)
	}
	return Tempt
}

func main() {
	A := "\\phi^2"
	fmt.Println(A)
	B := Powize(A)
	fmt.Println(B)
}
