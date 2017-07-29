package main

import (
	"fmt"

	"github.com/Axect/Go/CS/prob4"
)

func main() {
	c := make(chan float64)
	for i := 0; i < 10; i++ {
		go func(chan float64) {
			c <- pi.Pi2D(1000000) // Transfer pi value to channel
		}(c)
	}

	sum := 0.

	for j := 0; j < 10; j++ {
		sum += <-c / 10. // Transfer value in channel to sum
	}
	close(c)
	fmt.Println(sum)
}
