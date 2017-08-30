package main

import (
	"fmt"

	"github.com/Axect/csv"
	"github.com/wcharczuk/go-chart"
)

func main() {
	txt := csv.Read("/home/kavis/Documents/Go/src/github.com/Axect/RGE/Data/170_0_50.csv")
	array := Convert(txt)
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues:
			}
		}
	}
}

func Convert(str [][]string) [][]float64 {
	Temp := make([][]float64, len(str), len(str))
	for i := range str {
		for j := range str[i] {
			Temp[i][j] = strconv.ParseFloat(str[i][j])
		}
	}
	return Temp
}
