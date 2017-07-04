package slice

import (
	"fmt"
)

// ExampleSlice is example of slice
func ExampleSlice() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
}

// ExampleAppend is example of append
func ExampleAppend() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
}

// CopySlice is copy int slice
func CopySlice(x []int) []int {
	temp := make([]int, len(x))
	// Auto enumerate
	for i, elem := range x {
		temp[i] = elem
	}
	return temp
}
