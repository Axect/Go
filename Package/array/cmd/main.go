package main

import (
	"fmt"
	"reflect"

	"github.com/Axect/Go/Package/array"
)

func main() {
	A := array.Create(1, 1, 100)
	fmt.Println(A)
	fmt.Println(reflect.TypeOf(A))
	B := array.Float2Int(A)
	fmt.Println(B)
	fmt.Println(reflect.TypeOf(B))
}
