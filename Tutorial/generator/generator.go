package generator

// NewIntGenerator returns function which returns number of call
func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}
