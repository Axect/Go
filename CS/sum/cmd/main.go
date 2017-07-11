package main

import (
	"fmt"

	"github.com/Axect/Go/CS/sum"
)

func main() {
	A := sum.MakeSeq()
	resultA, timeA := sum.SummingList(A) // type(timeA) = time.Duration

	B, T := sum.MakeSeqPointer()
	resultB, timeB := sum.SummingPointer(B)

	n := 0
	for i := 1; i <= 9; i++ {
		_, time1 := sum.SummingList(A)
		_, time2 := sum.SummingPointer(B)
		timeA += time1
		timeB += time2
		n = i + 1
	}

	NewtimeA := timeA / 10
	NewtimeB := timeB / 10

	fmt.Println("Time of Make Array: ", T)
	fmt.Println("-Sum Module")
	fmt.Printf("Results: %v, Elapsed Time: %v (mean for %v times)\n", resultA, NewtimeA, n)
	fmt.Println("-Sum Pointer")
	fmt.Printf("Results: %v, Elapsed Time: %v (mean for %v times)\n", resultB, NewtimeB, n)
}
