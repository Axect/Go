package Polynomial

// Diff differentiates polynomial
func Diff(P Polynomial) Polynomial {
	L := len(P.Coeff)
	var Q Polynomial
	Q.Coeff = make(Vector, L-1, L-1)
	l := len(Q.Coeff)
	for i := range Q.Coeff {
		Q.Coeff[i] = float64(l-i) * P.Coeff[i]
	}
	return Q
}
