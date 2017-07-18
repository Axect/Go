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
)

var Initial = Coordinate{-9.851920196143998e-01, 1.316466809434336e-01, -4.877392224782687e-06}
var Initial2 = Coordinate{-9.864337701483683e-01, 1.230799243164879e-01, -4.374019384763304e-06}

// RunCoord runs Coordinate
func (C *Coordinate) RunCoord(V Velocity) {
	C.x += V.x
	C.y += V.y
	C.z += V.z
}

func (C Coordinate) Norm() float64 {
	N := math.Sqrt(math.Pow(C.x, 2) + math.Pow(C.y, 2) + math.Pow(C.z, 2))
	return N
}

// RunVel runs VelocitY
func (V *Velocity) RunVel(A Acceleration) {
	V.x += A.x
	V.y += A.y
	V.z += A.z
}

// CalcAccel calculates Acceleration each point
func (A *Acceleration) CalcAccel(C Coordinate) {
	ac := G * M / math.Pow(C.Norm(), 2)
	A.x = ac * C.x / C.Norm()
	A.y = ac * C.y / C.Norm()
	A.z = ac * C.z / C.Norm()
}

func test() {
	var A Acceleration
	fmt.Println(A.CalcAccel(Initial))
}
