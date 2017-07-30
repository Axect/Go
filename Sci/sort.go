package main

import (
	"github.com/Axect/Go/Package/csv"
)

func main() {
	data := csv.Read("Data/Notice2.csv")
	for i, row := range data {
		if row[8] == "" && row[7] != "" {
			for j := i + 1; ; j++ {
				if data[j][0] != "" && data[j][1] == "" && data[j][2] == "" {
					row[7] += data[j][0]
				} else if data[j][2] == "" && data[j][1] != "" && data[j][0] == "" {
					row[7] += data[j][1]
				} else if data[j][0] == "" && data[j][1] != "" {
					for k := 8; k <= 27; k++ {
						row[k] = data[j][k-7]
					}
				} else if data[j][0] != "" && (data[j][1] == "science" || data[j][1] == "admin") {
					row[7] += data[j][0]
					for k := 8; k <= 27; k++ {
						row[k] = data[j][k-7]
					}
				} else {
					break
				}
			}
		}
	}
	csv.Write(data, "Data/Test.csv")
}

// func main() {
// 	data := csv.Read("Data/Test.csv")
// 	for i := range data {
// 		if data[i][0] == "" || data[i][1] == "" {
// 			for j := range data[i] {
// 				data[i][j] = ""
// 			}
// 		}
// 	}
// 	csv.Write(data, "Data/Test2.csv")
// }
