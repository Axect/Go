package prob2

import (
	"fmt"
)

// Date is default struct
type Date struct {
	year  int
	month int
	day   int
}

// Ref is reference for month
var Ref = map[int]int{
	1:  31,
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
func (d Date) Days() int {
	s := 0
	if d.year%4 == 0 {
		if d.year%100 != 0 || d.year%400 == 0 {
			Ref[2] = 29
			for i := 1; i < d.month; i++ {
				s += Ref[i]
			}
			s += d.day
		} else {
			Ref[2] = 28
			for i := 1; i < d.month; i++ {
				s += Ref[i]
			}
			s += d.day
		}
	} else {
		Ref[2] = 28
		for i := 1; i < d.month; i++ {
			s += Ref[i]
		}
		s += d.day
	}
	return s
}

// DoCalendar generates calendar
func DoCalendar() {
	fmt.Println("---------------------------------------")
	fmt.Println("Calendar")
	fmt.Println("---------------------------------------")
	d := Date{2017, 5, 31}
	s := 0
	for i := 1; i < d.year; i++ {
		x := Date{i, 12, 31}
		s += x.Days()
	}
	s += d.Days()
	fmt.Println(s)
}
