package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/Axect/Go/CS/prob4"
	"github.com/Axect/Go/Package/Time"
)

const (
	n = 1
)

func main() {
	runtime.GOMAXPROCS(4)
	c := make(chan float64)
	defer Time.TimeTrack(time.Now(), "Total Process")
	for i := 0; i < n; i++ {
		go func(chan float64) {
			c <- pi.Pi2D(10000000) // Transfer pi value to channel
		}(c)
	}

	sum := 0.

	for j := 0; j < n; j++ {
		sum += <-c / float64(n) // Transfer value in channel to sum
	}
	fmt.Printf("Approx pi: %v, Error: %v%%\n", sum, math.Abs((math.Pi-sum)/math.Pi*100))
}
