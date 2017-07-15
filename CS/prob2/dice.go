package prob2

import (
	"fmt"
	"math/rand"
	"time"
)

// Roll rolls dice
func Roll(n int) []int64 {
	Result := make([]int64, n, n)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range Result {
		results := rand.Intn(6) + 1
		Result[i] = int64(results)
	}
	return Result
}

// Quincunx returns probability of quincunx
func Quincunx(array []int64) float64 {
	length := float64(len(array))
	count := 0
	for _, elem := range array {
		if elem == 5 {
			count++
		}
	}
	return float64(count) / length
}

// =============================================================================
// Roll Dice!
// =============================================================================

// DoRoll do roll
func DoRoll() {
	fmt.Println("---------------------------------------")
	fmt.Println("Roll Dice")
	fmt.Println("---------------------------------------")
	A := Roll(1e+06)
	P := Quincunx(A)
	fmt.Printf("The Probability of Quincunx is %v\n", P)
	fmt.Printf("where 1/6 is %v\n", 1./6.)
}
