package main

import (
	"runtime"
	"time"

	"github.com/Axect/Go/CS/prob2"
)

func main() {
	runtime.GOMAXPROCS(4)
	go prob2.DoRoll()
	go prob2.DoSort()
	time.Sleep(100 * time.Millisecond)
	prob2.DoRSP()
	prob2.DoCalendar()
}
