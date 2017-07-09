package sum

import (
	"time"

	"github.com/gonum/floats"
)

// MakeSeqPointer makes 1e+06 List Pointer
func MakeSeqPointer() *[1e+6]int {
	var seq [1e+6]int
	Seq := &seq
	for i := range Seq {
		Seq[i] = i + 1
	}
	return &seq
}

// MakeSeq makes 1e+06 Slice
func MakeSeq() []float64 {
	seq := make([]float64, 1e+6, 1e+6)
	for i := range seq {
		seq[i] = float64(i + 1)
	}
	return seq
}

//SummingPointer is sum sequence (pointer) return time
func SummingPointer(seq *[1e+6]int) (int, time.Duration) {
	s := 0
	start := time.Now()
	for _, elem := range seq {
		s += elem
	}
	elapsed := time.Since(start)
	return s, elapsed
}

//SummingList is sum list
func SummingList(seq []float64) (float64, time.Duration) {
	start := time.Now()
	s := floats.Sum(seq)
	elapsed := time.Since(start)
	return s, elapsed
}
