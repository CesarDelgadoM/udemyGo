package main

import "fmt"

type person struct {
	name string
	last string
}

type agente struct {
	person
	lpm bool
}

type human interface {
	presentarse()
}

func main() {
	x, y := saludar("Cesar", "Delgado")
	fmt.Println(x, y)
	sum1 := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sum1)
	b := []int{121, 45, 67, 123, 768}
	sum2 := sum(b...) //Pasar cada uno de los elementos del slice
	fmt.Println(sum2)
	defer foo() //difiere la funcion a que se ejecute al final siempre
	bar()
	crearAgente()
	polimorfismo()
	funcAnonimas()
	expFuncion()
	fmt.Println(returnFunc()())
	fmt.Println("El total de pares es:", sumPares(sum, b...))
	closures()
	fmt.Println(factorial(5))
}

func woo(st string) string {
	return fmt.Sprint("Hola desde woo, ", st)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, ` dice "hola".`)
	q := true
	return p, q
}

//Debe quedar (...) siempre al final de los parametros
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, n := range x {
		sum = sum + n
	}
	return sum
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func crearAgente() {
	a1 := agente{
		person: person{
			name: "cesar",
			last: "delgado",
		},
	}
	a1.presentarse()
}

func polimorfismo() {
	ag := agente{
		person: person{
			name: "Paola",
			last: "Avella",
		},
		lpm: true,
	}
	p := person{
		name: "Jose",
		last: "Medina",
	}
	ag.presentarse()
	p.presentarse()
	poli(ag)
	poli(p)
}

func (ag agente) presentarse() {
	fmt.Println("Hola soy el agente", ag.name, ag.last)
}

func (p person) presentarse() {
	fmt.Println("Hola soy la persona", p.name, p.last)
}

func poli(h human) {
	//Afirmacion de tipos
	switch h.(type) {
	case person:
		fmt.Println("(person)", h.(person).name)
	case agente:
		fmt.Println("(agente)", h.(agente).name)
	}
	fmt.Println("Estoy siendo comportandome como:", h)
	fmt.Printf("%T\n", h)
}

// Funciones anonimas
func funcAnonimas() {
	func() {
		fmt.Println("La funcion anonima ase ejecuto")
	}()
	func(num int) {
		fmt.Println("El numero de la suerte es:", num)
	}(33)
}

//Expresion funcion
func expFuncion() {
	f := func() {
		fmt.Println("Mi primera exprecion funcion")
	}
	g := func(year int) {
		fmt.Println("tu año de nacimiento es:", year)
	}
	f()
	g(1996)
}

//Retorno de una funcion
func returnFunc() func() int {
	return func() int {
		return 1996
	}
}

//CallBack
func sumPares(f func(x ...int) int, nums ...int) int {
	var y []int
	for _, v := range nums {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

func closures() {
	a := incrementar()
	b := incrementar()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementar() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

//Recursividad
func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
