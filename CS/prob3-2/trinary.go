package trinary

import (
	"fmt"
	"log"
)

type Trinary struct {
	value int
}

func (T Trinary) String() string {
	return fmt.Sprintf("%d [Trinary]\n", T.value)
}

func (T *Trinary) Assign(x int) {
	T.value = x % 3
}

func AS(S, T Trinary, b bool) Trinary {
	R := Trinary{0}
	if b {
		R.Assign(S.value + T.value)
	} else {
		R.Assign(Abs(S.value - T.value))
	}
	return R
}

func Mul(S, T Trinary) Trinary {
	R := Trinary{0}
	R.Assign(S.value * T.value)
	return R
}

func (T *Trinary) Square() {
	T.Assign(T.value * T.value)
}

func Input() Trinary {
	var T Trinary
	var x int
	fmt.Print("Please input a positive integer: ")
	_, err := fmt.Scanln(&x)
	if err != nil || x < 0 {
		log.Fatal("Please input a positive integer!")
	}
	T.Assign(x)
	return T
}

func Test() {
	S := Input()
	T := Input()
	R := AS(S, T, true)
	R.Square()
	fmt.Println(R)
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
}
