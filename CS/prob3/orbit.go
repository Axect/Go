package orbit

import (
	"fmt"
	"math"
	"time"

	"github.com/Axect/Go/Package/csv"
)

// Coordinate handle cartesian coordinate
type Coordinate struct {
	x float64
	y float64
	z float64
}

// Velocity handle velocity vector
type Velocity Coordinate

// Acceleration handle acceleration vector
type Acceleration Coordinate

const (
	G     = 6.67259e-11       // Gravitational Constant
	m     = 5.9736e+24        // Earth mass
	M     = 1.9891e+30        // Sun mass
	AU    = 1.49597870691e+11 // Astronomy Unit
	tstep = 43200             // Time Step
	N     = 730 * 10
)

var Initial = Coordinate{-9.851920196143998e-01 * AU, 1.316466809434336e-01 * AU, -4.877392224782687e-06 * AU}
var Initial2 = Coordinate{-9.864337701483683e-01 * AU, 1.230799243164879e-01 * AU, -4.374019384763304e-06 * AU}

var TestArray [2][N + 1]float64

var CList [N + 1]Coordinate
var TList [N + 1]float64
var EList [N + 1]float64

func (C Coordinate) String() string {
	return fmt.Sprintf("x:%v, y:%v, z:%v", C.x/AU, C.y/AU, C.z/AU)
}

func (C *Coordinate) Add(D Coordinate) {
	C.x += D.x
	C.y += D.y
	C.z += D.z
}

func (C *Coordinate) Sub(D Coordinate) {
	C.x -= D.x
	C.y -= D.y
	C.z -= D.z
}

func (C *Coordinate) Div(t float64) {
	C.x = C.x / t
	C.y = C.y / t
	C.z = C.z / t
}

func (C Coordinate) Mul(t float64) Coordinate {
	var D Coordinate
	D.x = C.x * t
	D.y = C.y * t
	D.z = C.z * t
	return D
}

func (V *Velocity) CalcVel(C, D Coordinate) {
	C.Sub(D)     // C - D
	C.Div(tstep) // C / tstep
	V.x = C.x    // V = (C-D)/tstep
	V.y = C.y
	V.z = C.z
}

func (V Velocity) String() string {
	return fmt.Sprintf("x:%v, y:%v, z:%v", V.x/AU, V.y/AU, V.z/AU)
}

func (A Acceleration) String() string {
	return fmt.Sprintf("x:%v, y:%v, z:%v", A.x/AU, A.y/AU, A.z/AU)
}

func (A Acceleration) Mul(t float64) Acceleration {
	var B Acceleration
	B.x = A.x * t
	B.y = A.y * t
	B.z = A.z * t
	return B
}

// RunCoord runs Coordinate
func (C *Coordinate) RunCoord(V Velocity, A Acceleration) {
	C.x += V.x*tstep + A.x*math.Pow(tstep, 2)/2
	C.y += V.y*tstep + A.y*math.Pow(tstep, 2)/2
	C.z += V.z*tstep + A.z*math.Pow(tstep, 2)/2
}

func (C Coordinate) Norm() float64 {
	n := math.Sqrt(math.Pow(C.x, 2) + math.Pow(C.y, 2) + math.Pow(C.z, 2))
	return n
}

func (V Velocity) Norm() float64 {
	n := math.Sqrt(math.Pow(V.x, 2) + math.Pow(V.y, 2) + math.Pow(V.z, 2))
	return n
}

func Calc(C [N + 1]Coordinate) [N + 1]float64 {
	E := &EList
	var V Velocity
	V.CalcVel(Initial2, Initial)
	E[0] = math.Pow(V.Norm(), 2)*m/2 - G*M*m/Initial.Norm()
	for i, co := range C {
		if i >= 1 {
			temp := co.Mul(1.) // Generate new Coordinate
			temp.Sub(C[i-1])   // C[i] - C[i-1]
			temp.Div(tstep)    // (C[i] - C[i-1]) / tstep
			n := math.Pow(temp.Norm(), 2)
			E[i] = n*m/2 - G*M*m/co.Norm()
		}
	}
	return EList
}

// RunVel runs VelocitY
func (V *Velocity) RunVel(A Acceleration) {
	V.x += A.x * tstep
	V.y += A.y * tstep
	V.z += A.z * tstep
}

// CalcAccel calculates Acceleration each point
func (A *Acceleration) CalcAccel(C Coordinate) {
	ac := -1. * G * M / math.Pow(C.Norm(), 3)
	A.x = ac * C.x
	A.y = ac * C.y
	A.z = ac * C.z
}

// NIntegration do numercal integration (Euler Method)
func NIntegration() (Coordinate, Velocity, Acceleration, float64) {
	X := Initial
	T := 0.
	Pest := &TestArray[0]
	Qest := &TestArray[1]
	Pest[0] = X.x / AU
	Qest[0] = T
	var V Velocity
	A := Acceleration{0, 0, 0}
	V.CalcVel(Initial2, Initial)
	X.RunCoord(V, A)
	T += tstep
	Pest[1] = X.x / AU
	Qest[1] = T / 86400
	for i := 2; i <= N; i++ {
		A.CalcAccel(X)
		V.RunVel(A)
		X.RunCoord(V, A)
		T += tstep
		Pest[i] = X.x / AU
		Qest[i] = T / 86400
	}
	return X, V, A, T / 86400
}

// Verlet is numerical integration algorithm for MD
func Verlet(I Coordinate) ([N + 1]Coordinate, [N + 1]float64) {
	C := I
	T := 0.
	var V Velocity
	TL := &TList
	CL := &CList
	CL[0] = C
	TL[0] = T
	A := Acceleration{0, 0, 0}
	V.CalcVel(Initial2, Initial)
	C.RunCoord(V, A)
	T += tstep
	CL[1] = C
	TL[1] = T / 86400
	for i := 2; i <= N; i++ {
		A.CalcAccel(CL[i-1])
		temp := CList[i-1].Mul(2)               // 2*x_n
		temp.Sub(CList[i-2])                    // 2*x_n - x_(n-1)
		B := A.Mul(math.Pow(float64(tstep), 2)) // A*(h^2)
		Temp := Coordinate{B.x, B.y, B.z}
		temp.Add(Temp) // 2*x_n - x_(n-1) + A*h^2
		CL[i] = temp
		T += tstep
		TL[i] = T / 86400
	}
	return CList, TList
}

func Test() {
	WriteArray := make([][]string, N+1, N+1)
	WriteArrayE := make([][]string, N+1, N+1)
	start := time.Now()
	//X, V, A, T := NIntegration()
	C, T := Verlet(Initial)
	E := Calc(C)
	elapsed := time.Since(start)
	for i := range C {
		WriteArray[i] = []string{fmt.Sprint(C[i].x / AU), fmt.Sprint(T[i])}
		WriteArrayE[i] = []string{fmt.Sprint(E[i]), fmt.Sprint(T[i])}
	}
	//for i := range C {
	//	WriteArray[i] = []string{fmt.Sprint(E[i]), fmt.Sprint(T[i])}
	//}
	//csv.Write(WriteArray, "../Data/test_taylor.csv")
	csv.Write(WriteArray, "../Data/X_Ver.csv")
	csv.Write(WriteArrayE, "../Data/Energy.csv")
	//fmt.Printf("%v,\n %v,\n %v,\n %v,\n", X, V, A, T)
	fmt.Println("time is ", elapsed)
}
