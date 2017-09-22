package Polynomial

// Add adds two polynomials
func Add(P, Q Polynomial) Polynomial {
	L1, L2 := len(P.Coeff), len(Q.Coeff)
	var R Polynomial
	switch {
	case L1 > L2:
		R.Coeff = make([]float64, L1, L1)
		for i := 0; i < L1; i++ {
			if i < L1-L2 {
				R.Coeff[i] = P.Coeff[i]
			} else {
				j := i - (L1 - L2)
				R.Coeff[i] = P.Coeff[i] + Q.Coeff[j]
			}
		}
	case L1 < L2:
		R.Coeff = make([]float64, L2, L2)
		for i := 0; i < L2; i++ {
			if i < L2-L1 {
				R.Coeff[i] = Q.Coeff[i]
			} else {
				j := i - (L1 - L2)
				R.Coeff[i] = P.Coeff[j] + Q.Coeff[i]
			}
		}
	default:
		R.Coeff = make([]float64, L1, L1)
		for i := 0; i < L1; i++ {
			R.Coeff[i] = P.Coeff[i] + Q.Coeff[i]
		}
	}
	return R
}

// Sub subtracts two polynomials
func Sub(P, Q Polynomial) Polynomial {
	R := Minus(Q)
	result := Add(P, R)
	return result
}
