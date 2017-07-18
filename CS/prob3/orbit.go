package orbit

import (
	"fmt"
	"math"
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
	Gf    = G / (AU * AU)
)

var Initial = Coordinate{-9.851920196143998e-01, 1.316466809434336e-01, -4.877392224782687e-06}
var Initial2 = Coordinate{-9.864337701483683e-01, 1.230799243164879e-01, -4.374019384763304e-06}

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

func (V *Velocity) CalcVel(C, D Coordinate) {
	C.Sub(D)
	C.Div(tstep)
	V.x = C.x
	V.y = C.y
	V.z = C.z
}

// RunCoord runs Coordinate
func (C *Coordinate) RunCoord(V Velocity) {
	C.x += V.x * tstep
	C.y += V.y * tstep
	C.z += V.z * tstep
}

func (C Coordinate) Norm() float64 {
	N := math.Sqrt(math.Pow(C.x, 2) + math.Pow(C.y, 2) + math.Pow(C.z, 2))
	return N
}

// RunVel runs VelocitY
func (V *Velocity) RunVel(A Acceleration) {
	V.x += A.x * tstep
	V.y += A.y * tstep
	V.z += A.z * tstep
}

// CalcAccel calculates Acceleration each point
func (A *Acceleration) CalcAccel(C Coordinate) {
	ac := Gf * M / math.Pow(C.Norm(), 2)
	A.x = ac * C.x / C.Norm()
	A.y = ac * C.y / C.Norm()
	A.z = ac * C.z / C.Norm()
}

func Test() {
	//var A Acceleration
	var V Velocity
	X := Initial
	V.CalcVel(Initial2, Initial)
	X.RunCoord(V)
	fmt.Println(X)
}
