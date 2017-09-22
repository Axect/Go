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

// Integrate integrates Polynomial (Caution : You should input const)
func Integrate(P Polynomial, C float64) Polynomial {
	L := len(P.Coeff)
	var Q Polynomial
	Q.Coeff = make(Vector, L+1, L+1)
	l := len(Q.Coeff) - 1
	for i := range Q.Coeff {
		switch {
		case i == l:
			Q.Coeff[i] = C
		default:
			Q.Coeff[i] = P.Coeff[i] / float64(l-i)
		}
	}
	return Q
}
