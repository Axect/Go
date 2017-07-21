package test

import (
	"fmt"
	"math"
)

const (
	Mp       = 1.221 * 1E+19
	MpR      = 2.4 * 1E+18 // Reduced Planck Mass
	MW       = 80.385
	MZ       = 91.1876
	MH       = 125.09
	alphasMZ = 0.1182
)

type RGE struct {
	t   float64
	lH  float64
	yt  float64
	g1  float64
	g2  float64
	g3  float64
	phi float64
}

type Action interface {
	Initailize()
	Running()
}

func (R *RGE) Initialize(mt, xi float64) {
	R.t = 0
	R.lH = 0.12604 + 0.00206*(MH-125.15) - 0.00004*(mt-173.34)
	R.yt = (0.93690 + 0.00556*(mt-173.34) - 0.00003*(MH-125.15) - 0.00042*(alphasMZ-0.1184)/0.0007)
	R.g3 = 1.1666 + 0.00314*(alphasMZ-0.1184)/0.007 - 0.00046*(mt-173.34)
	R.g2 = 0.64779 + 0.00004*(mt-173.34) + 0.00011*(MW-80.384)/0.014
	R.g1 = 0.35830 + 0.00011*(mt-173.34) - 0.00020*(MW-80.384)/0.014
}

func (R *RGE) Running() {
	h := math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
	R.phi = h
	sh := (1 + 1)
}

func Test() {
	var R RGE
	mt, xi := 170.85, 50.
	R.Initialize(mt, xi)
	fmt.Println(R.lH)
}
