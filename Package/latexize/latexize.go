package latexize

import (
	"fmt"
	"strings"
)

// Powize makes TeX code to Go Code
func Powize(n string) string {
	Tempt := n
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
		Tempt = fmt.Sprintf("math.Pow(%v,%v)", Front, Back)
	}
	return Tempt
}

// LaTeXize means LaTeX code to Go code
func LaTeXize(sentence string) string {
	words := strings.Split(sentence, " ")
	output := ""
	for _, word := range words {
		output += Powize(word) + " "
	}
	return output
}
