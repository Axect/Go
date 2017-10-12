package main

import (
	sr "github.com/Axect/Go/Advanced/Lorentz"
)

func main() {
	A := sr.NewCovector([]float64{1., 2., 3., 4.})
	A.Show()
}
