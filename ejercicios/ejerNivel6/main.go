package main

import (
	"fmt"
	"math"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {
	fmt.Println("Mi nombre es", p.nombre, "y mi edad es", p.edad)
}

type cuadrado struct {
	l int
	h int
}

type circulo struct {
	r float64
}

func (s cuadrado) getArea() float64 {
	return float64(s.l * s.h)
}

func (c circulo) getArea() float64 {
	return math.Pi * c.r * c.r
}

type forma interface {
	getArea() float64
}

func info(f forma) {
	switch f.(type) {
	case cuadrado:
		fmt.Println("el area del cuadrado es:", f.(cuadrado).getArea())
	case circulo:
		fmt.Println("el area del circulo es:", f.(circulo).getArea())
	}

}

func main() {
	exercise4()
	exercise5()
	exercise6()
	exercise7()
}

func exercise4() {
	p := persona{
		nombre:   "Paola",
		apellido: "Avella",
		edad:     27,
	}
	p.presentar()
}

func exercise5() {
	c := circulo{
		r: 12.345,
	}
	s := cuadrado{
		l: 15,
		h: 15,
	}
	info(c)
	info(s)
}

func exercise6() {
	fib := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func fibo() func() int {
	a := 1
	b := 1
	sum := 0
	return func() int {
		sum = a + b
		a = b
		b = sum
		return sum
	}
}

func exercise7() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(revisarLen, x))
}

func revisarLen(v ...int) int {
	lenV := len(v)
	if lenV == 0 {
		return 0
	} else if lenV == 1 {
		return v[0]
	} else {
		return v[0] + v[lenV-1]
	}
}

func foo(f func(v ...int) int, v []int) int {
	x := f(v...)
	x++
	return x
}
