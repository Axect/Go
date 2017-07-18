package main

import (
	"fmt"

	"github.com/Axect/Go/slice"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(slice.CopySlice(x))
}
