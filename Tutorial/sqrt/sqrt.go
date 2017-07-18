package sqrt

import (
	"fmt"
	"math"
)

// Sqrt is Sqrt
func Sqrt(x float64) string {
	if x >= 0 {
		return fmt.Sprint(math.Sqrt(x))
	}
	return Sqrt(-x) + "i"
}
