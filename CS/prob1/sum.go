package prob1

import (
	"fmt"
	"time"

	"github.com/gonum/floats"
)

// MakeSeqPointer makes 1e+06 List Pointer
func MakeSeqPointer() (*[1e+6]int, time.Duration) {
	start := time.Now()
	var seq [1e+6]int
	Seq := &seq
	for i := range Seq {
		Seq[i] = i + 1
	}
	elapsed := time.Since(start)
	return &seq, elapsed
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

// DoSum is Do Sum
func DoSum() {
	A := MakeSeq()
	resultA, timeA := SummingList(A) // type(timeA) = time.Duration

	B, T := MakeSeqPointer()
	resultB, timeB := SummingPointer(B)

	n := 0
	for i := 1; i <= 9; i++ {
		_, time1 := SummingList(A)
		_, time2 := SummingPointer(B)
		timeA += time1
		timeB += time2
		n = i + 1
	}

	NewtimeA := timeA / 10
	NewtimeB := timeB / 10

	fmt.Println("---------------------------------------------")
	fmt.Println("Effective Sum List")
	fmt.Println("---------------------------------------------")
	fmt.Println("Time of Make Array: ", T)
	fmt.Println("-Sum Module")
	fmt.Printf("Results: %v, Elapsed Time: %v (mean for %v times)\n", resultA, NewtimeA, n)
	fmt.Println("-Sum Pointer")
	fmt.Printf("Results: %v, Elapsed Time: %v (mean for %v times)\n", resultB, NewtimeB, n)
	fmt.Println()
}
