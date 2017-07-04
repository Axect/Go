package array

import (
	"fmt"
)

// Array is just study array
func Array() {
	example := [3]int{1, 2, 3}
	for i, elem := range example {
		fmt.Println(i, elem)
	}
}
