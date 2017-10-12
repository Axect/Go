package Lorentz

// Special makes normal object to relativistic object.
// It maps int to VectorType
var Special = map[TensorType]string{
	0: "Scalar",
	1: "Covariant Vector",
	2: "Contravariant Vector",
	3: "Tensor",
}
