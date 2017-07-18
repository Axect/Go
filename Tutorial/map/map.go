package main

import "fmt"

var dict = map[string]string{
	"Proton":  "uud",
	"Neutron": "udd",
}

func main() {
	for keys, values := range dict {
		fmt.Println("Key:", keys, "Value:", values)
	}
}
