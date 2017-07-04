package main

import (
	"fmt"
)

func add(x, y string) string {
	return x + y
}

func main() {
	fmt.Println("Happy", add("Hello", "Hi"), "Day")
}
