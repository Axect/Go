package main

import (
	"fmt"
	"time"

	"github.com/Axect/Go/Tutorial/generator2"
)

func main() {
	start := time.Now()
	gen := generator2.IntGenerator()
	for i := 1; i < 100; i++ {
		gen()
	}
	fmt.Println(gen())
	elapsed := time.Since(start)

	start2 := time.Now()
	s := 0
	for i := 0; i <= 100; i++ {
		s = i
	}
	fmt.Println(s)
	elapsed2 := time.Since(start2)
	fmt.Println("Generator:", elapsed, "For Loop:", elapsed2)
}
