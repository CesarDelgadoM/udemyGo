package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func cambiame(p *persona) {
	p2 := persona{
		nombre:   "Paola",
		apellido: "Avella",
		edad:     27,
	}
	*p = p2
	fmt.Println(*p)
	(*p).nombre = "Jose"
	fmt.Println(*p)
}

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	a := "Hello World"
	var p *string = &a
	var p2 **string = &p
	fmt.Println(&p, p, *p)
	fmt.Println(&p2, p2, *p2)
	**p2 = "Hola mundo"
	fmt.Println(&p, p, *p)
	fmt.Println(&p2, p2, **p2)
}

func exercise2() {
	p := persona{
		nombre:   "Cesar",
		apellido: "Delgado",
		edad:     25,
	}
	cambiame(&p)
}
