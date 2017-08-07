package matrix

// Tridiagonal is a method to solve A[i]X[i-1] + B[i]X[i] + C[i]X[i+1] = D[i] wrt X[i]
func Tridiagonal(A, B, C, D []float64) []float64 {
	P := make([]float64, len(A), len(A))
	Q := make([]float64, len(A), len(A))
	for i := 0; i < len(A); i++ {
		if i == 0 {
			P[i] = C[i] / B[i]
			Q[i] = D[i] / B[i]
		} else {
			P[i] = C[i] / (B[i] - A[i]*P[i-1])
			Q[i] = (D[i] - A[i]*Q[i-1]) / (B[i] - A[i]*P[i-1])
		}
	}
	X := make([]float64, len(A), len(A))
	for i := len(A) - 1; i >= 0; i-- {
		if i == len(A)-1 {
			X[i] = Q[i]
		} else {
			X[i] = Q[i] - P[i]*X[i+1]
		}
	}
	return X
}
