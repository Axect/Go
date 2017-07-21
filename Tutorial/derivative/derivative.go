package derivative

// Normal for convenience (Type is great!)
type Normal func(float64) float64

// Derivative make function to Derivative
func Derivative(f Normal) Normal {
	h := 1e-05
	return func(x float64) float64 {
		return (f(x+h) - f(x-h)) / (2 * h)
	}
}
