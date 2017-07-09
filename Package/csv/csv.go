package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Write writes csv file
func Write(List [][]string, name string) {
	Title := name + ".csv"
	file, err := os.Create(Title)
	checkError("Cannot create a file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range List {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	fmt.Print("Complete to Write")
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
