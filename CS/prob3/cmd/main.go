package main

import (
	"fmt"
	"runtime"

	"github.com/Axect/Go/CS/prob3"
)

func main() {
	runtime.GOMAXPROCS(8)
	go orbit.DoOrbit()
	fmt.Scanln()
}
