package main

import (
	"encoding/csv"
	"log"
	"os"
)

// Hello test
var Hello = [][]string{{"Hello"}, {"World!"}}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot Create File", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range Hello {
		err := writer.Write(value)
		checkError("Cannot Write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
