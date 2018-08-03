package main

import "fmt"

func main() {
	var A int
	n, msg := fmt.Scanln(&A)
	fmt.Println(n, A)
	fmt.Println(msg) // Error message
}
