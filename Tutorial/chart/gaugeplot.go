package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Axect/csv"
	"github.com/wcharczuk/go-chart"
)

const (
	DefaultDPI = 300
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	//func drawChart() {
	txt := csv.Read("/home/kavis/Documents/Go/src/github.com/Axect/RGE/Data/Gauge_170_0_50.csv")
	array := csv.Convert(txt)
	var T, lH, g1, g2, g3 []float64
	for _, List := range array {
		T = append(T, List[0])
		lH = append(lH, List[1])
		g1 = append(g1, List[3])
		g2 = append(g2, List[4])
		g3 = append(g3, List[5])
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "t",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Gauge",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		DPI:    100,
		Width:  1000,
		Height: 600,
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "g1",
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0),
				},
				XValues: T,
				YValues: g1,
			},
			chart.ContinuousSeries{
				Name: "g2",
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(1),
				},
				XValues: T,
				YValues: g2,
			},
			chart.ContinuousSeries{
				Name: "g3",
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(2),
				},
				XValues: T,
				YValues: g3,
			},
		},
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	res.Header().Set("Content-Type", "image/png")
	//Render(RendererProvider, io.Writer)
	err := graph.Render(chart.PNG, res)
	if err != nil {
		log.Fatal("Why Don't you Know")
	}

	//collector := &chart.ImageWriter{}
	//graph.Render(chart.PNG, collector)
	//image, err := collector.Image()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Final Image: %dx%d\n", image.Bounds().Size().X, image.Bounds().Size().Y)
	//collector.Write(image)
	fmt.Println("Complete!")
}

func main() {
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//drawChart()
}
