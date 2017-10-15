package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%*d ", 3, 2*i)
	}
	fmt.Println()
	for i := 1; i < 10; i++ {
		fmt.Printf("%*d ", 3, 8*i)
	}
}
