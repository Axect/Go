package generator2

// IntGenerator is generator which generates new int
func IntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}
