package prob1

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
	// Generates error to scan not int value or negative
	if _, err := fmt.Scanln(&n); err != nil || n < 0 {
		fmt.Println("Error: Integer n should be positive or zero")
		os.Exit(1)
	}
	m := float64(n) + 0.5
	repr := fmt.Sprint(m) + "hbar"
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

//DoHarm is DO Harm
func DoHarm() {
	fmt.Println("---------------------------------------------")
	fmt.Println("Harmonic Oscillator")
	fmt.Println("---------------------------------------------")
	Estr := Harmonic()
	Ereal := Converter(Estr)
	fmt.Println("Energy (symbolic): ", Estr)
	fmt.Println("Energy (numeric): ", Ereal)
	fmt.Println()
}
