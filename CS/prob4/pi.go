package pi

import (
	"math"
	"math/rand"
	"time"

	"github.com/Axect/Go/Package/Time"
)

// Coord2D represent 2D Coordinate
type Coord2D struct {
	x float64
	y float64
}

type Coord3D struct {
	x float64
	y float64
	z float64
}

// Radius is method for Coord2D to calculate radius
func (C *Coord2D) Radius() float64 {
	return math.Pow(C.x, 2) + math.Pow(C.y, 2)
}

func (D *Coord3D) Radius() float64 {
	return math.Pow(D.x, 2) + math.Pow(D.y, 2) + math.Pow(D.z, 2)
}

// Pi2D is function for calculate pi
func Pi2D(n int) float64 {
	defer Time.TimeTrack(time.Now(), "Pi2D")
	rand.Seed(time.Now().UTC().UnixNano())
	count := 0.
	for i := 0; i < n; i++ {
		C := Coord2D{rand.Float64(), rand.Float64()}
		if C.Radius() <= 1 {
			count++
		}
	}
	return count / float64(n) * 4
}

// Pi3D is function for calculate pi in 3D
func Pi3D(n int) float64 {
	defer Time.TimeTrack(time.Now(), "Pi3D")
	rand.Seed(time.Now().UTC().UnixNano())
	count := 0.
	for i := 0; i < n; i++ {
		C := Coord3D{rand.Float64(), rand.Float64(), rand.Float64()}
		if C.Radius() <= 1 {
			count++
		}
	}
	return count / float64(n) * 6
}
