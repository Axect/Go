package prob2

import (
	"fmt"
	"os"
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

// Ref2 for days
var Ref2 = map[int]string{
	0: "Sunday",
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
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

// DayTotal obtain days from 0001.01.01
func (d Date) DayTotal() int {
	s := 0
	for i := 1; i < d.year; i++ {
		x := Date{i, 12, 31}
		s += x.Days()
	}
	s += d.Days()
	return s
}

// DayFinder generates day
func DayFinder(n int) string {
	x := n % 7
	return Ref2[x]
}

// DoCalendar generates calendar
func DoCalendar() {
	fmt.Println("---------------------------------------")
	fmt.Println("Calendar")
	fmt.Println("---------------------------------------")
	var year, month, day int
	fmt.Print("날짜를 입력하세요 (년, 월, 일): ")
	_, err := fmt.Scanln(&year, &month, &day)
	if err != nil {
		fmt.Println("값이 잘못되었습니다.")
		os.Exit(1)
	}
	d := Date{year, month, day}
	fmt.Printf("입력하신 날짜는 서기 %d일이며 %s입니다.\n", d.DayTotal(), DayFinder(d.DayTotal()))
}
