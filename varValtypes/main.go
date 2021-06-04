package main

import "fmt"

var a int
var b string = "Programa"
var c bool
var d bool = true

//Creamos un tipo de dato propio
type dinero int

var din dinero

func main() {
	example1()
	example2()
	example3()
}

func example1() {
	e := 32
	f := "dice hola mundo"
	g := `El doctor dice que comer vegetales
			es saludable`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(a, b, c, d, e, f, g)
}

func example2() {
	//Representacion del valor con el verbo %v
	fmt.Printf("El valor de la variable a es %v\n", a)
	fmt.Printf("El valor de la variable b es %v\n", b)
	//Representacion del tipo con el verbo %T
	fmt.Printf("El tipo de la variable a es %T\n", a)
	fmt.Printf("El tipo de la variable b es %T\n", b)
	/*
		The default format:

		bool:                     %t
		int, int8, etc.:          %d
		uint, uint8, etc.:        %d, %#x if printed with %#v
		float32, complex64, etc.: %g
		string:                   %s
		chan:                     %p
		pointer:                  %p
	*/
	fmt.Printf("El valor de la variable a es %d\n", a)
	fmt.Printf("El valor de la variable b es %s\n", b)

	s1 := fmt.Sprint("El ", b, " dice hola mundo")
	fmt.Println(s1)
}

func example3() {
	din = 10000
	fmt.Println(din)
	fmt.Printf("%T", din)
}
