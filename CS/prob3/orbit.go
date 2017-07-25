package orbit

import (
	"fmt"
	"math"
	"time"

	"github.com/Axect/Go/Package/Time"
	"github.com/Axect/Go/Package/csv"
)

const (
	G     = 6.67259e-11       // Gravitational Constant
	m     = 5.9736e+24        // Earth mass
	M     = 1.9891e+30        // Sun mass
	AU    = 1.49597870691e+11 // Astronomy Unit
	tstep = 43200             // Time Step
	N     = 730 * 10
)

type Vector struct {
	x, y, z float64
}

type VList struct {
	R [N + 1]Vector
}

type VPair struct {
	r Vector
	v Vector
}

func Initialize() (Vector, Vector) {
	A := Vector{-9.851920196143998e-01 * AU, 1.316466809434336e-01 * AU, -4.877392224782687e-06 * AU}
	B := Vector{-9.864337701483683e-01 * AU, 1.230799243164879e-01 * AU, -4.374019384763304e-06 * AU}
	return A, B
}

func (v Vector) String() string {
	return fmt.Sprintf("%v, %v, %v", v.x, v.y, v.z)
}

func Plus(v, w Vector) Vector {
	var A Vector
	A.x = v.x + w.x
	A.y = v.y + w.y
	A.z = v.z + w.z
	return A
}

func (c *Vector) Add(w Vector) {
	c.x += w.x
	c.y += w.y
	c.z += w.z
}

func Mul(v Vector, f float64) Vector {
	var A Vector
	A.x = v.x * f
	A.y = v.y * f
	A.z = v.z * f
	return A
}

func (V *VList) Assign(v Vector, i int) {
	V.R[i] = v
}

func (v Vector) Norm() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2))
}

func En(c, v Vector) (float64, float64) {
	T := 1. / 2. * m * math.Pow(v.Norm(), 2)
	U := (-1. * G * M * m) / c.Norm()
	return T, U
}

func (P *VPair) Run() {
	a := Mul(P.r, -1.*G*M/math.Pow(P.r.Norm(), 3))
	P.v.Add(Mul(a, tstep))
	P.r.Add(Mul(P.v, tstep))
}

func Taylor(v1, v2 Vector) (VList, [N + 1]float64, [N + 1]float64) {
	var C VList
	var T, U [N + 1]float64
	defer Time.TimeTrack(time.Now(), "Taylor")
	t, u := &T, &U
	i1, i2 := v1, v2
	v0 := Plus(Mul(v2, 1./tstep), Mul(v1, -1./tstep))
	P, NULL := VPair{i2, v0}, Vector{0., 0., 0.}
	C.Assign(i1, 0)
	C.Assign(i2, 1)
	_, u[0] = En(i1, NULL)
	_, u[1] = En(i2, NULL)
	for i := 2; i <= N; i++ {
		vf := P.v
		P.Run()
		vm := Plus(Mul(vf, 0.5), Mul(P.v, 0.5))
		if i == 2 {
			vd := Plus(vm, Mul(vf, -1.))
			t[0], _ = En(NULL, Plus(vf, Mul(vd, -1.)))
		}
		if i == N {
			vd := Plus(P.v, Mul(vf, -1.))
			t[N], _ = En(NULL, Plus(P.v, vd))
		}
		C.Assign(P.r, i)
		t[i-1], u[i] = En(P.r, vm)
	}
	return C, T, U
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

func ConvertF(F, G [N + 1]float64) [][]string {
	W := make([][]string, N+1, N+1)
	for j := 0; j <= N; j++ {
		W[j] = []string{fmt.Sprint(F[j]), fmt.Sprint(G[j])}
	}
	return W
}

func DoOrbit() {
	C, T1, U1 := Taylor(Initialize())
	D, T2, U2 := Taylor(Reversize(C))
	ERR := Plus(D.R[N], Mul(C.R[0], -1.))
	fmt.Printf(" Number of years: %v\n Errors of Coordinates(AU): %v\n", N/730, Mul(ERR, 1./AU))
	W1 := Convert(C)
	W2 := Convert(D)
	W3 := ConvertF(T1, U1)
	W4 := ConvertF(T2, U2)
	csv.Write(W1, "Data/taylor.csv")
	csv.Write(W2, "Data/reverse.csv")
	csv.Write(W3, "Data/energy.csv")
	csv.Write(W4, "Data/energy_rev.csv")
}
