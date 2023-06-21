package main

import "fmt"

func main() {
	strings := []string{"hello", "world", "Geovanny"}
	// i => indice
	// s => elemento
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
