package Lorentz

// CoVector is covariant vector
type CoVector struct {
	Vector //anonymous field vector
}

// ContraVector is contravariant vector
type ContraVector struct {
	Vector
}

// NewCovector makes covector
func NewCovector(vec []float64) CoVector {
	V := Vector{
		Comp: vec,
		Type: 1,
	}
	return CoVector{Vector: V}
}

// NewContraVector makes contravector
func NewContraVector(vec []float64) ContraVector {
	V := Vector{
		Comp: vec,
		Type: 2,
	}
	return ContraVector{Vector: V}
}
