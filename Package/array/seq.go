package array

// Vector is Type Alias in Go 1.9
type Vector = []float64

// Create generates arithmetic Sequence
func Create(init, step, end float64) Vector {
	Length := int((end - init + 1.) / step)
	A := make(Vector, Length, Length)
	s := init
	for i := range A {
		switch i {
		case 0:
			A[i] = s
		default:
			s += step
			A[i] = s
		}
	}
	return A
}
