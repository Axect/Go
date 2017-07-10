package main

import (
	"fmt"
	"time"

	"github.com/Axect/Go/Tutorial/method"
)

func main() {
	A := method.Person{
		Name:    "Axect",
		Account: 1,
	}
	start := time.Now()
	A.Add()
	elapsed := time.Since(start)
	fmt.Println(A.Account, elapsed)
}
