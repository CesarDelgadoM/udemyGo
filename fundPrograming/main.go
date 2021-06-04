package main

import (
	"fmt"
	"runtime"
)

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
}

func example1() {
	var i int8 = 20
	fmt.Println(i)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func example2() {
	s1 := "Hola mundo"
	s2 := `esta es una linea
			partida`
	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n", s1)
	fmt.Println(s2)
	fmt.Printf("El tipo de s1 es: %T\n", s2)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("El tipo de s1 es: %T\n", bs)
}

func example3() {
	s1 := "Hello World"
	tam := len(s1)
	for i := 0; i < tam; i++ {
		//Unicode code point
		fmt.Printf("%#U \n", s1[i])
	}

	for i, l := range s1 {
		fmt.Printf("En el indice %d el valor es %#x\n", i, l)
	}

	fmt.Printf("Con el verbo %q indico que se imprima el string %s\n", "%s", s1)
}

const a = 42
const b = 42.32
const c = "Cesar Delgado"

//iota un incrementador para cada vez que aparece la palabra const
const (
	p = iota
	y
	z
)

const (
	t = iota
	r
	e
)

type name string

var x name

func example4() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	x = c
	fmt.Println(x)

	fmt.Println(p)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(t)
	fmt.Println(r)
	fmt.Println(e)
}

//Bit shifting
const (
	_  = iota
	kb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func example5() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)
}
