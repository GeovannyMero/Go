package main

import "fmt"

/*
	Un puntero continene la direccion de memoria de un valor

	El tipo *T es un puntero a un valor de T. Su valor cero es nulo.

	El operador & genera un puntero a su operador

	El operador * denota el valor subyacente del puntero
*/

func main() {
	i, j := 42, 201

	p := &i         //puntero a i
	fmt.Println(*p) //leer  i a traves del puntero
	*p = 21         //asigna i a traves del puntero
	fmt.Println(i)  //imprime el nuevo valor de i

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
