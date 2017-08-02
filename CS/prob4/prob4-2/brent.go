package root

import (
	"fmt"
	"log"
	"math"
)

// CheckBisect checks bisection
func CheckBisect(f Normal, a, b float64) bool {
	if f(a)*f(b) >= 0 {
		return false
	}
	return true
}

// Brent is brent method
func Brent(f Normal, a, b float64) float64 {
	if !CheckBisect(f, a, b) {
		log.Fatal("Can't find root in this interval")
	}
	if math.Abs(f(a)) < math.Abs(f(b)) {
		a, b = b, a
	}
	c := a
	var s, d float64
	mflag := true
	for f(s) == 0 || math.Abs(b-a) < 1e-17 {
		if f(a) != f(c) && f(b) != f(c) {
			s = a*f(b)*f(c)/((f(a)-f(b))*(f(a)-f(c))) + b*f(a)*f(c)/((f(b)-f(a))*(f(b)-f(c))) + c*f(a)*f(b)/((f(c)-f(a))*(f(c)-f(b))) // Inverse Qudratic Interpolation
		} else {
			s = b - f(b)*(b-a)/(f(b)-f(a)) // Secant Method
		}
		switch {
		case s <= (3*a+b)/4 || s >= b:
			s = (a + b) / 2
			mflag = true
		case mflag && math.Abs(s-b) >= math.Abs(b-c)/2:
			s = (a + b) / 2
			mflag = true
		case !mflag && math.Abs(s-b) >= math.Abs(c-d)/2:
			s = (a + b) / 2
			mflag = true
		case mflag && math.Abs(b-c) < 1e-16:
			s = (a + b) / 2
			mflag = true
		case !mflag && math.Abs(c-d) < 1e-16:
			s = (a + b) / 2
			mflag = true
		default:
			mflag = false
		}
		d = c
		c = b
		if f(a)*f(s) < 0 {
			b = s
		} else {
			a = s
		}
		if math.Abs(f(a)) < math.Abs(f(b)) {
			a, b = b, a
		}
		fmt.Println(s)
	}
	return s
}
