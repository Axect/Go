package test

import (
	"fmt"
	"math"

	"github.com/Axect/Go/Package/csv"
)

const (
	Mp       = 1.221 * 1E+19
	MpR      = 2.4 * 1E+18 // Reduced Planck Mass
	MW       = 80.385
	MZ       = 91.1876
	MH       = 125.09
	alphasMZ = 0.1182
	h        = 1e-04
	Step     = 1e+04 * 44
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

type Container [Step]float64

type Action interface {
	Initailize()
	Running()
}

type Gamma func(float64, float64) float64

func (R *RGE) Initialize(mt, xi float64) {
	R.t = 0.
	R.lH = 0.12604 + 0.00206*(MH-125.15) - 0.00004*(mt-173.34)
	R.yt = (0.93690 + 0.00556*(mt-173.34) - 0.00003*(MH-125.15) - 0.00042*(alphasMZ-0.1184)/0.0007)
	R.g3 = 1.1666 + 0.00314*(alphasMZ-0.1184)/0.007 - 0.00046*(mt-173.34)
	R.g2 = 0.64779 + 0.00004*(mt-173.34) + 0.00011*(MW-80.384)/0.014
	R.g1 = 0.35830 + 0.00011*(mt-173.34) - 0.00020*(MW-80.384)/0.014
}

func (R *RGE) Running(mt, xi float64) {
	hg := math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
	R.phi = hg
	sh := (1 + xi*math.Pow(hg, 2)) / math.Pow(hg, 2) / (1 + (1+6*xi)*xi*math.Pow(hg, 2)/math.Pow(MpR, 2))
	BlH1 := 6*(1+3*math.Pow(sh, 2))*math.Pow(R.lH, 2) + 12*R.lH*math.Pow(R.yt, 2) - 6*math.Pow(R.yt, 4) - 3*R.lH*(3*math.Pow(R.g2, 2)+math.Pow(R.g1, 2)) + 3/8*(2*math.Pow(R.g2, 4)+math.Pow((math.Pow(R.g1, 2)+math.Pow(R.g2, 2)), 2))
	Bg11 := (81 + sh) / 12 * math.Pow(R.g1, 3)
	Bg21 := -(39 - sh) / 12 * math.Pow(R.g2, 3)
	Bg31 := -7 * math.Pow(R.g3, 3)
	Byt1 := R.yt * ((23/6+2/3*sh)*math.Pow(R.yt, 2) - (8*math.Pow(R.g3, 2) + 9/4*math.Pow(R.g2, 2) + 17/12*math.Pow(R.g1, 2)))
	gamma1 := -1 / (16 * math.Pow(math.Pi, 2)) * (9/4*math.Pow(R.g2, 2) + 3/4*math.Pow(R.g1, 2) - 3*math.Pow(R.yt, 2))

	//Calc Beta function
	g := MakeBeta(gamma1)
	Bg1, Bg2, Bg3 := g(Bg11, 0.), g(Bg21, 0.), g(Bg31, 0.)
	BlH, Byt := g(BlH1, 0.), g(Byt1, 0.)

	// Real Running
	R.lH += h * BlH
	R.yt += h * Byt
	R.g1 += h * Bg1
	R.g2 += h * Bg2
	R.g3 += h * Bg3
	R.t += h
}

func Convert(C1, C2, C3 Container) [][]string {
	W := make([][]string, Step, Step)
	for i := range C1 {
		W[i] = []string{fmt.Sprint(C1[i]), fmt.Sprint(C2[i]), fmt.Sprint(C3[i])}
	}
	return W
}

func Run() {
	var R RGE
	var g1, g2, g3 Container
	q1, q2, q3 := &g1, &g2, &g3
	mt, xi := 170.85, 50.
	R.Initialize(mt, xi)
	q1[0], q2[0], q3[0] = R.g1, R.g2, R.g3
	for i := 1; i < Step; i++ {
		R.Running(mt, xi)
		q1[i], q2[i], q3[i] = R.g1, R.g2, R.g3
	}
	W := Convert(g1, g2, g3)
	csv.Write(W, "Data/gauge.csv")
}

func MakeBeta(g float64) Gamma {
	return func(f1, f2 float64) float64 {
		temp := 1./(16*math.Pow(math.Pi, 2))*f1 + math.Pow(1./(16*math.Pow(math.Pi, 2)), 2)*f2
		return temp / (1 + g)
	}
}
