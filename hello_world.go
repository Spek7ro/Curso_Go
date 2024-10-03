package main

// paquete de formateo
import (
	"fmt"
)

func main() {
	fmt.Println("Hello World! GO")

	/*
		Comentarios de varias lineas
	*/

	// Variables
	// Se puede inferir el tipo de dato
	var myString string = "Mi String"
	fmt.Println(myString)

	myString = "Mi String Modificada"
	fmt.Println(myString)

	/*
		Si la variable se define con un tipo de dato,
		no se puede cambiar el tipo de dato de la variable
	*/

	// Por defecto un int es 32 bits
	var myInt int = 10
	fmt.Println("Mi int es:", myInt)
	fmt.Println("Mi int + 1:", myInt+1)
	fmt.Println("Mi int:", myInt)

	// Para un numero de 64 bits
	var myInt64 int64 = 1000000000000000000
	fmt.Println("Mi int64 es:", myInt64)
	fmt.Printf("%s %d", myString, myInt)

}
