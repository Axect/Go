package main

import (
	. "github.com/Axect/Go/Package/Numeric/Lagrange"
	"github.com/Axect/Go/Package/array"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func main() {
	X := []float64{1., 2., 3., 4.}
	Y := []float64{1., 4., 9., 16.}
	L := LPolynomial(X, Y)
	T := array.Arithmetic(0., 0.01, 400)
	Q := make([]float64, len(T), len(T))
	for i, t := range T {
		Q[i] = L(t)
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Lagrange Plots"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(T))
	for i := range pts {
		pts[i].X = T[i]
		pts[i].Y = Q[i]
	}

	plotutil.AddLines(p,
		"Interp", pts,
	)

	if err := p.Save(10*vg.Inch, 8*vg.Inch, "Fig/Lag2.svg"); err != nil {
		panic(err)
	}
	// csv.Write(temp, "Data/lag.csv")
}
