package main

import (
	"fmt"

	tn "github.com/chewxy/gorgonia/tensor"
)

func main() {
	A := tn.New(tn.WithShape(3, 3), tn.WithBacking([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(A)
	fmt.Println(A.Shape())
}
