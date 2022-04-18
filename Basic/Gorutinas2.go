package main

import (
	"fmt"
	"time"
)

func saludar(nombres []string) {
	for _, nombre := range nombres {
		fmt.Printf("Hola %s\n", nombre)
		time.Sleep(time.Millisecond * 500)
	}
}

func despsdir(nombres []string) {
	for _, nombre := range nombres {
		fmt.Printf("Adios %s\n", nombre)
		time.Sleep(time.Second)

	}
}

func main() {
	nombres := []string{"Orlando", "Gissell", "Mateo"}
	go saludar(nombres)
	go despsdir(nombres)

	var s string
	fmt.Scan(&s)
}
