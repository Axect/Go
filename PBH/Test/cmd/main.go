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
	wg.Add(20)
	// For Laptop
	runtime.GOMAXPROCS(20)
	// For Server
	// runtime.GOMAXPROCS(20)
	start := time.Now()
	for i := 1; i <= 20; i++ {
		go func(n int) {
			title := fmt.Sprintf("Gauge%d.csv", n)
			test.Run(170.8+0.1*float64(n), 100, title)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Total time is ", elapsed)
}
