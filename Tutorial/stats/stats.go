package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {
	data := []float64{1, 2, 3, 4, 4, 5}

	cov, _ := stats.Covariance(data, data)
	median, _ := stats.Median(data)
	fmt.Println(median)
	fmt.Println(cov)

	a := stats.Coordinate{1., 1.}
	b := stats.Coordinate{2., 2.}
	d := stats.Coordinate{3., 2.5}
	c := stats.Series{a, b, d}
	L, _ := stats.LinearRegression(c)
	fmt.Println(L)
}
