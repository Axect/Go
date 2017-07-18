package main

import (
	"fmt"

	"github.com/Axect/Go/Tutorial/generator"
)

func main() {
	gen := generator.NewIntGenerator()
	fmt.Println(gen(), gen(), gen())
}
