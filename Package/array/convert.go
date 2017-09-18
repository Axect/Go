package array

// Float2Int converts []float64 to []int64
func Float2Int(A []float64) []int64 {
	B := make([]int64, len(A), len(A))
	for i, elem := range A {
		B[i] = int64(elem)
	}
	return B
}
