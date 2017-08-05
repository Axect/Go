package product

// Product products sequence from init to end
func Product(seq []float64) float64 {
	result := 1.
	for i := 0; i < len(seq); i++ {
		result *= seq[i]
	}
	return result
}
