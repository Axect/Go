package array

// Arithmetic generates arithmetic Sequence
func Arithmetic(init, d float64, end int) []float64 {
	A := make([]float64, end, end)
	s := init
	for i := range A {
		switch i {
		case 0:
			A[i] = s
		default:
			s += d
			A[i] = s
		}
	}
	return A
}
