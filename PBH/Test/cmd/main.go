package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/Axect/Go/PBH/Test"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// For Laptop
	runtime.GOMAXPROCS(2)
	// For Server
	// runtime.GOMAXPROCS(20)
	start := time.Now()
	go func() {
		test.Run(170.85, 50., "Gauge50.csv")
		defer wg.Done()
	}()
	go func() {
		test.Run(170.85, 100., "Gauge100.csv")
		defer wg.Done()
	}()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Total time is ", elapsed)
}
