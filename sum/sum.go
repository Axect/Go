package main

import "fmt"

// Sum55 is summing 1 ~ n
func Sum55(n int) int {
	S := 0
	for i := 1; i <= n; i++ {
		S += i
	}
	return S
}

func main() {
	fmt.Println(Sum55(10))
}
