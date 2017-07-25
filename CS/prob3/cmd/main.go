package main

import (
	"runtime"

	"github.com/Axect/Go/CS/prob3"
)

func main() {
	runtime.GOMAXPROCS(20)
	orbit.DoOrbit()
}
