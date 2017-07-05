package main

import (
	"fmt"
	"time"

	"github.com/Axect/Go/Tutorial/pointer"
)

var array [1e6]int
var slice = make([]int, 1e6)

func main() {
	start := time.Now()
	fmt.Println(pointer.SumList(array))
	elapsed := time.Since(start)

	start2 := time.Now()
	fmt.Println(pointer.SumPointer(&array))
	elapsed2 := time.Since(start2)

	start3 := time.Now()
	fmt.Println(pointer.SumSlice(slice))
	elapsed3 := time.Since(start3)

	fmt.Println("Array:", elapsed, "Pointer:", elapsed2, "Slice:", elapsed3)
}
