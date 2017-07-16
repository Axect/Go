package prob2

import (
	"fmt"
)

type Date struct {
	year  int
	month int
	day   int
}

// Ref is reference for month
var Ref = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

// Days generates number of days
func (d Date) Days() {
	s := 0
	if d.year%4 == 0 && d.year%100 != 0 {
		for i := 1; i < d.month; i++ {
			s += Ref[i]
		}
		s += d.day
	}
	fmt.Println(s)
}
