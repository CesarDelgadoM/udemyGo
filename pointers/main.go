package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type forma interface {
	area() float64
}

func (c *circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func info(f forma) {
	fmt.Println("Area:", f.area())
}

func main() {
	example1()
	example2()
}

func example1() {
	x := 33
	fmt.Println(x)  //valor
	fmt.Println(&x) //direccion de memoria

	fmt.Printf("%T\n", &x)

	var y *int = &x
	fmt.Println(y)
	fmt.Println(*y)

	*y = 56
	fmt.Println(*y)

	//Paso por valor
	a := 1
	pasoValor(a)
	fmt.Println(a)
	//Paso por referencia
	pasoReferencia(&a)
	fmt.Println(a)
}

func pasoValor(x int) {
	x++
	fmt.Println(x)
}

func pasoReferencia(x *int) {
	*x++
	fmt.Println(*x)
}

func example2() {
	c := circulo{
		radio: 33,
	}
	info(&c)
}
