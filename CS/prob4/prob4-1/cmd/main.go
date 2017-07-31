package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/Axect/Go/CS/prob4/prob4-1"
	"github.com/Axect/Go/Package/Time"
)

const (
	n = 1
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan float64)
	d := make(chan float64)
	defer Time.TimeTrack(time.Now(), "Total Process")
	for i := 0; i < n; i++ {
		go func(chan float64) {
			c <- pi.Pi2D(10000000) // Transfer pi value to channel
		}(c)
		go func(chan float64) {
			d <- pi.Pi3D(10000000) // Transfer pi value to channel
		}(d)
	}

	sum1 := 0.
	sum2 := 0.

	for j := 0; j < n; j++ {
		sum1 += <-c / float64(n) // Transfer value in channel to sum
		sum2 += <-d / float64(n)
	}
	fmt.Printf("Approx pi 2D: %v, Error: %v%%\n", sum1, math.Abs((math.Pi-sum1)/math.Pi*100))
	fmt.Printf("Approx pi 3D: %v, Error: %v%%\n", sum2, math.Abs((math.Pi-sum2)/math.Pi*100))
}
