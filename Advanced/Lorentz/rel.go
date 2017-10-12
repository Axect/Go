package Lorentz

// CoVector is covariant vector
type CoVector struct {
	Vector //anonymous field vector
}

func NewCovector(vec []float64) CoVector {
	V := Vector{
		Comp: vec,
		Type: "Covariant Vector",
	}
	return CoVector{Vector: V}
}
