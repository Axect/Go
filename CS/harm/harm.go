package harm

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// hReal is dirac constant in unit (eV)
	hReal = 6.582119514 * 1e-16
)

// Harmonic returns energy in string
func Harmonic() string {
	var n int
	fmt.Print("Input n:")
	fmt.Scanln(&n)
	if n < 0 {
		fmt.Println("Error: n should be positive or zero")
		os.Exit(1)
	}
	m := float64(n) + 0.5
	repr := fmt.Sprint(m+1/2) + "hbar"
	return repr
}

// Converter converts string expr to float and finally Sprint
func Converter(h string) string {
	hl := len(h)
	h = h[:hl-4]
	r, _ := strconv.ParseFloat(h, 64)
	r = r * hReal
	return fmt.Sprintf("%veV", r)
}
