package test

import (
	"fmt"
	"math"
	"time"

	"github.com/Axect/Go/Package/csv"
)

const (
	Mp       = 1.221 * 1E+19 // Plank Mass
	MpR      = 2.4 * 1E+18   // Reduced Planck Mass
	MW       = 80.385        // Mass of W
	MZ       = 91.1876       // Mass of Z
	MH       = 125.09        // Mass of Higgs
	alphasMZ = 0.1182        // alphas(MZ)
	h        = 1e-04         // precision
	Step     = 1e+04 * 44    // Number of lists
)

// RGE : To RGE Running
type RGE struct {
	t   float64
	lH  float64
	yt  float64
	g1  float64
	g2  float64
	g3  float64
	phi float64
}

// Container is container of coupling constants
type Container [Step]float64

// Bay is []container
type Bay []Container

// Action represents all methods
type Action interface {
	Initailize()
	Running()
}

// Gamma for convenience
type Gamma func(float64, float64) float64

// Initialize is initialize by some constants
func (R *RGE) Initialize(mt, xi float64) {
	R.t = 0.
	R.lH = 0.12604 + 0.00206*(MH-125.15) - 0.00004*(mt-173.34)
	R.yt = (0.93690 + 0.00556*(mt-173.34) - 0.00003*(MH-125.15) - 0.00042*(alphasMZ-0.1184)/0.0007)
	R.g3 = 1.1666 + 0.00314*(alphasMZ-0.1184)/0.007 - 0.00046*(mt-173.34)
	R.g2 = 0.64779 + 0.00004*(mt-173.34) + 0.00011*(MW-80.384)/0.014
	R.g1 = 0.35830 + 0.00011*(mt-173.34) - 0.00020*(MW-80.384)/0.014
}

// Running is main action
func (R *RGE) Running(mt, xi float64) {
	hg := math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
	R.phi = hg
	sh := (1 + xi*math.Pow(hg, 2)) / math.Pow(MpR, 2) / (1 + (1+6*xi)*xi*math.Pow(hg, 2)/math.Pow(MpR, 2))

	// 1-loop Beta function
	BlH1 := 6*(1+3*math.Pow(sh, 2))*math.Pow(R.lH, 2) + 12*R.lH*math.Pow(R.yt, 2) - 6*math.Pow(R.yt, 4) - 3*R.lH*(3*math.Pow(R.g2, 2)+math.Pow(R.g1, 2)) + 3./8*(2*math.Pow(R.g2, 4)+math.Pow((math.Pow(R.g1, 2)+math.Pow(R.g2, 2)), 2))
	Bg11 := (81 + sh) / 12 * math.Pow(R.g1, 3)
	Bg21 := -(39 - sh) / 12 * math.Pow(R.g2, 3)
	Bg31 := -7 * math.Pow(R.g3, 3)
	Byt1 := R.yt * ((23./6+2./3*sh)*math.Pow(R.yt, 2) - (8*math.Pow(R.g3, 2) + 9./4*math.Pow(R.g2, 2) + 17./12*math.Pow(R.g1, 2)))
	gamma1 := -1. / (16 * math.Pow(math.Pi, 2)) * (9./4*math.Pow(R.g2, 2) + 3./4*math.Pow(R.g1, 2) - 3*math.Pow(R.yt, 2))

	// 2-loop Beta function
	BlH2 := 1./48*((912+3*sh)*math.Pow(R.g2, 6)-(290-sh)*math.Pow(R.g1, 2)*math.Pow(R.g2, 4)-(560-sh)*math.Pow(R.g1, 4)*math.Pow(R.g2, 2)-(380-sh)*math.Pow(R.g1, 6)) + (38-8*sh)*math.Pow(R.yt, 6) - math.Pow(R.yt, 4)*(8./3*math.Pow(R.g1, 2)+32*math.Pow(R.g3, 2)+(12-117*sh+108*math.Pow(sh, 2))*R.lH) + R.lH*(-1./8*(181+54*sh-162*math.Pow(sh, 2))*math.Pow(R.g2, 4)+1./4*(3-18*sh+54*math.Pow(sh, 2))*math.Pow(R.g1, 2)*math.Pow(R.g2, 2)+1./24*(90+377*sh+162*math.Pow(sh, 2))*math.Pow(R.g1, 4)+(27+54*sh+27*math.Pow(sh, 2))*math.Pow(R.g2, 2)*R.lH+(9+18*sh+9*math.Pow(sh, 2))*math.Pow(R.g1, 2)*R.lH-(48+288*sh-324*math.Pow(sh, 2)+624*math.Pow(sh, 3)-324*math.Pow(sh, 4))*math.Pow(R.lH, 2)) + math.Pow(R.yt, 2)*(-9./4*math.Pow(R.g2, 4)+21./2*math.Pow(R.g1, 2)*math.Pow(R.g2, 2)-19./4*math.Pow(R.g1, 4)+R.lH*(45./2*math.Pow(R.g2, 2)+85./6*math.Pow(R.g1, 2)+80*math.Pow(R.g3, 2)-(36+108*math.Pow(sh, 2))*R.lH))
	Bg12 := 199./18*math.Pow(R.g1, 5) + 9./2*math.Pow(R.g1, 3)*math.Pow(R.g2, 2) + 44./3*math.Pow(R.g1, 3)*math.Pow(R.g3, 2) - 17./6*sh*math.Pow(R.g1, 3)*math.Pow(R.yt, 2)
	Bg22 := 3./2*math.Pow(R.g1, 2)*math.Pow(R.g2, 3) + 35./6*math.Pow(R.g2, 5) + 12*math.Pow(R.g2, 3)*math.Pow(R.g3, 2) - 3./2*sh*math.Pow(R.g2, 3)*math.Pow(R.yt, 2)
	Bg32 := 11./6*math.Pow(R.g1, 2)*math.Pow(R.g3, 3) + 9./2*math.Pow(R.g2, 2)*math.Pow(R.g3, 3) - 26*math.Pow(R.g3, 5) - 2*sh*math.Pow(R.g3, 3)*math.Pow(R.yt, 2)
	Byt2 := R.yt * (-23./4*math.Pow(R.g2, 4) - 3./4*math.Pow(R.g1, 2)*math.Pow(R.g2, 2) + 1187./216*math.Pow(R.g1, 4) + 9*math.Pow(R.g2, 2)*math.Pow(R.g3, 2) + 19./9*math.Pow(R.g1, 2)*math.Pow(R.g3, 2) - 108*math.Pow(R.g3, 4) + (225./16*math.Pow(R.g2, 2)+131./16*math.Pow(R.g1, 2)+36*math.Pow(R.g3, 2))*sh*math.Pow(R.yt, 2) + 6*(-2*math.Pow(sh, 2)*math.Pow(R.yt, 4)-2*math.Pow(sh, 3)*math.Pow(R.yt, 2)*R.lH+math.Pow(sh, 2)*math.Pow(R.lH, 2)))
	gamma2 := 1. / math.Pow(16*math.Pow(math.Pi, 2), 2) * (271./32*math.Pow(R.g2, 4) - 9./16*math.Pow(R.g1, 2)*math.Pow(R.g2, 2) - 431./96*sh*math.Pow(R.g1, 4) - 5./2*(9./4*math.Pow(R.g2, 2)+17./12*math.Pow(R.g1, 2)+8*math.Pow(R.g3, 2))*math.Pow(R.yt, 2) + 27./4*sh*math.Pow(R.yt, 4) - 6*math.Pow(sh, 3)*math.Pow(R.lH, 2))

	//Calc Beta function
	g := MakeBeta(gamma1 + gamma2)
	Bg1, Bg2, Bg3 := g(Bg11, Bg12), g(Bg21, Bg22), g(Bg31, Bg32)
	BlH, Byt := g(BlH1, BlH2), g(Byt1, Byt2)

	// Real Running
	R.lH += h * BlH
	R.yt += h * Byt
	R.g1 += h * Bg1
	R.g2 += h * Bg2
	R.g3 += h * Bg3
	R.t += h
}

// Convert to write
func Convert(B Bay) [][]string {
	l := len(B)
	W := make([][]string, Step, Step)
	for i := range B[0] {
		W[i] = []string{fmt.Sprint(B[0][i])}
		for j := 1; j < l; j++ {
			W[i] = append(W[i], fmt.Sprint(B[j][i]))
		}
	}
	return W
}

// Run is main function
func Run(mt, xi float64, name string) {
	var R RGE
	var lH, g1, g2, g3, yt Container
	qH, q1, q2, q3, qt := &lH, &g1, &g2, &g3, &yt
	start := time.Now()
	R.Initialize(mt, xi)
	qH[0], q1[0], q2[0], q3[0], qt[0] = R.lH, R.g1, R.g2, R.g3, R.yt
	for i := 1; i < Step; i++ {
		R.Running(mt, xi)
		qH[i], q1[i], q2[i], q3[i], qt[i] = R.lH, R.g1, R.g2, R.g3, R.yt
	}
	elapsed1 := time.Since(start)
	start2 := time.Now()
	B := Bay{lH, g1, g2, g3, yt}
	W := Convert(B)
	title := fmt.Sprintf("./Data/%s", name)
	csv.Write(W, title)
	elapsed2 := time.Since(start2)
	fmt.Printf("Running time is %v\nWriting time is %v\n", elapsed1, elapsed2)
}

// MakeBeta : Input gamma -> Output Beta function
func MakeBeta(g float64) Gamma {
	return func(f1, f2 float64) float64 {
		temp := 1./(16*math.Pow(math.Pi, 2))*f1 + math.Pow(1./(16*math.Pow(math.Pi, 2)), 2)*f2
		return temp / (1 + g)
	}
}
