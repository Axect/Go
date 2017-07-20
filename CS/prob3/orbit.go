package orbit

import (
	"fmt"
	"math"
	"time"
)

const (
	G     = 6.67259e-11       // Gravitational Constant
	m     = 5.9736e+24        // Earth mass
	M     = 1.9891e+30        // Sun mass
	AU    = 1.49597870691e+11 // Astronomy Unit
	tstep = 43200             // Time Step
	N     = 730 * 10
)

func Initialize() (Vector, Vector) {
	A := Vector{-9.851920196143998e-01 * AU, 1.316466809434336e-01 * AU, -4.877392224782687e-06 * AU}
	B := Vector{-9.864337701483683e-01 * AU, 1.230799243164879e-01 * AU, -4.374019384763304e-06 * AU}
	return A, B
}

type VList struct {
	R [N + 1]Vector
}

type Vector struct {
	x, y, z float64
}

func (v Vector) String() string {
	return fmt.Sprintf("%v, %v, %v", v.x, v.y, v.z)
}

func AS(v, w Vector, b bool) Vector {
	var q Vector
	if b {
		q.x, q.y, q.z = v.x+w.x, v.y+w.y, v.z+w.z
	} else {
		q.x, q.y, q.z = v.x-w.x, v.y-w.y, v.z-w.z
	}
	return q
}

func (v Vector) Mul(f float64) Vector {
	q := Vector{0, 0, 0}
	q.x, q.y, q.z = v.x*f, v.y*f, v.z*f
	return q
}

func (V *VList) Assign(v Vector, i int) {
	V.R[i] = v
}

func (v Vector) NormPow() (float64, float64) {
	f := math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2)
	g := math.Sqrt(f)
	return g, f
}

func Accel(coord Vector) Vector {
	n, _ := coord.NormPow()
	n = math.Pow(n, 3)
	ac := -1. * G * M / n
	a := coord.Mul(ac)
	return a
}

func Taylor(v1, v2 Vector) VList {
	var c, v, a Vector
	var C VList
	i1, i2 := v1, v2
	v0 := (AS(i2, i1, false)).Mul(1. / tstep)
	c, v = i1, v0
	C.Assign(c, 0)
	c = AS(c, v.Mul(tstep), true)
	C.Assign(c, 1)
	for i := 2; i <= N; i++ {
		a = Accel(c)
		v = AS(v, a.Mul(tstep), true)
		c = AS(c, v.Mul(tstep), true)
		C.Assign(c, i)
	}
	return C
}

func Reversize(V VList) (Vector, Vector) {
	f1, f2 := V.R[N], V.R[N-1]
	return f1, f2
}

func Convert(V VList) [][]string {
	W := make([][]string, N+1, N+1)
	for j := 0; j <= N; j++ {
		W[j] = []string{fmt.Sprint((V.R[j]).x / AU), fmt.Sprint((V.R[j]).y / AU), fmt.Sprint((V.R[j]).z / AU)}
	}
	return W
}

func DoOrbit() {
	start := time.Now()
	C := Taylor(Initialize())
	D := Taylor(Reversize(C))
	ERR := AS(D.R[N], C.R[0], false)
	elapsed := time.Since(start)
	fmt.Printf(" Number of years: %v\n Elapsed Time: %v\n Errors of Coordinates(AU): %v\n", N/730, elapsed, ERR.Mul(1./AU))
	//W1 := Convert(C)
	//W2 := Convert(D)
	//csv.Write(W1, "Data/taylor.csv")
	//csv.Write(W2, "Data/reverse.csv")
}
