package main

import (
	"runtime"
	"sync"

	"github.com/Axect/Go/CS/prob3"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	runtime.GOMAXPROCS(20)
	go func() {
		orbit.DoOrbit()
		defer wg.Done()
	}()
	wg.Wait()
	orbit.DoOrbit()
}
