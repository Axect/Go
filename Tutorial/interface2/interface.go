package main

import (
	"fmt"
	"math"
)

// Geometry can include Sphere, Cube and anything which has volume() method.
// Similar with Inheritance
type Geometry interface {
	volume() float64
}

// Sphere is struct for sphere
type Sphere struct {
	radius float64
}

// Cube is struct for cube
type Cube struct {
	length float64
}

func main() {
	var S Sphere
	var C Cube

	S.radius = 3.
	C.length = 6.

	fmt.Println(S.volume(), C.volume())
	fmt.Println(TotalVolume(&S, &C))
}

func (S *Sphere) volume() float64 {
	return 4. / 3. * math.Pi * math.Pow(S.radius, 3.)
}

func (C *Cube) volume() float64 {
	return math.Pow(C.length, 3)
}

// TotalVolume calculate total volume
func TotalVolume(geom ...Geometry) float64 {
	var vol float64
	for _, s := range geom {
		vol += s.volume()
	}
	return vol
}
