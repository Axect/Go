package main

import "fmt"

var array = [5]int{0: 10, 1: 20}
var matrix = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

// If not const, then can't determine vector which size is l1
const (
	l1 = len(matrix)
	l2 = len(matrix[0])
)

var vector [l1]int

func main() {
	array[3] = 4
	fmt.Println(&array)
	fmt.Println(matrix)
	for i, group := range matrix {
		vector[i] = group[0]
	}
	vector[2] = 100
	fmt.Println(vector)
	fmt.Println(matrix)
}
