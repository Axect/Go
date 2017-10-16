package array

func Sum(V Vector) float64 {
	s := 0.
	for _, elem := range V {
		s += elem
	}
	return s
}