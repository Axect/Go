package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	Files := Load()
	Folders := PickFolder(Files)
	fmt.Println(Folders)
}

// Load loads file names
func Load() []string {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	l := len(files)
	Names := make([]string, l, l)
	for i, file := range files {
		Names[i] = file.Name()
	}
	return Names
}

// PickFolder checks Folder (exception: binary file)
func PickFolder(Files []string) []string {
	var Folders []string

	for _, file := range Files {
		Info, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}

		if Info.IsDir() {
			Folders = append(Folders, file)
		}
	}

	return Folders
}
