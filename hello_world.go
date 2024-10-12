package main

// paquete de formateo
import (
	"fmt"
	"reflect"
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
	fmt.Printf("%s %d", myString, myInt) // formato de cadenas
	fmt.Println()
	fmt.Println(myString, myInt)

	// Ver el tipo de dato de una variable con reflect
	fmt.Println(reflect.TypeOf(myString))
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(myInt64))

	// variable flotante (por defecto 64 bits)
	var myFloat float64 = 10.5
	fmt.Println("Mi float es:", myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	// transformar un numero int en float
	fmt.Println(myFloat + float64(myInt))

	// tipos de dato booleanos
	var myBool bool = true
	fmt.Println("Mi bool es:", myBool)
	myBool = false
	fmt.Println("Mi bool es:", myBool)

	// Operador := asignar valor a variable
	// el operador := sirve para la declaración y 
	// asignación inicial de una variable
	myString3 := "Mi String 3"
	fmt.Println(myString3)

			

}
