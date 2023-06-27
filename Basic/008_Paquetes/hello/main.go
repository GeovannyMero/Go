package main

import (
	"fmt"

	"github.com/GeovannyMero/Go/Basic/008_Paquetes/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("GEovanny")
	fmt.Println(message)
}
