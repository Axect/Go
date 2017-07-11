package main

import (
	"fmt"

	"github.com/Axect/Go/CS/prob1"
)

func main() {
	prob1.DoHarm()
	go prob1.DoDiff()
	go prob1.DoSum()
	fmt.Scanln()
}
