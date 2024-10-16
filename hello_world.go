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

	// Constantes
	// se usa la palabra reservada const
	const myConst = "Mi Constante"
	const myConst2 = 10
	fmt.Println("Constante 1:", myConst)
	fmt.Println("Constante 2:", myConst2)

	// Control de flujo
	myInt = 11
	if myInt == 10 {
		fmt.Println("myInt es 10")
	} else if myInt == 11 {
		fmt.Println("myInt es 11")
	} else {
		fmt.Println("myInt no es 10")
	}

	/*
		Operadores
		+ - % * / ^ & | != == < > <= >= && ||
		+= -= *= /= %= ^= &= |=
	*/

	
	// Array por defecto se inicializa con ceros
	var myArray [5]int 
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)
	
	var myArray2 [5]string = [5]string{"uno", "dos", "tres", "cuatro", "cinco"}
	fmt.Println(myArray2)

	fmt.Println(myArray2[0])


}
