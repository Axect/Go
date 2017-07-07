package pointer

// SumList is sum of list
func SumList(array [1e6]int) int {
	sum := 0
	for _, elem := range array {
		sum += elem
	}
	return sum
}

// SumPointer is sum of Pointer
func SumPointer(array *[1e6]int) int {
	sum := 0
	for _, elem := range array {
		sum += elem
	}
	return sum
}

// SumSlice is sum of Slice
func SumSlice(array []int) int {
	sum := 0
	for _, elem := range array {
		sum += elem
	}
	return sum
}
