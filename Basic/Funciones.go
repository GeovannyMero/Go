package main

import "fmt"

//funcion toma dos parametros y retorna un entero
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 50)) //ejecuta la funcion "add" y lo muestra por pantalla
}
