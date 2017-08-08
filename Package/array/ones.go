package array

// Ones generates ones array with float64
func Ones(length int) []float64 {
	A := make([]float64, length, length)
	for i := range A {
		A[i] = 1.
	}
	return A
}
