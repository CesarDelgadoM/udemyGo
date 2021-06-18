package main

import (
	"fmt"

	"github.com/CesarDelgadoM/udemyGo/ejercicios/ejerNivel13/matematicas"
	"github.com/CesarDelgadoM/udemyGo/ejercicios/ejerNivel13/palabra"
	"github.com/CesarDelgadoM/udemyGo/ejercicios/ejerNivel13/perrito"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	//exercise1()
	//exercise2()
	exercise3()
}

func exercise1() {
	fido := canino{
		nombre: "Fido",
		edad:   perrito.Edad(10),
	}
	fmt.Println(fido)
	fmt.Println(perrito.EdadDos(20))
}

func exercise2() {
	//fmt.Println(palabra.Conteo(cita.Cuando))
	for k, v := range palabra.ConteoUso("Hola Hola mundo") {
		fmt.Println(v, k)
	}
}

func exercise3() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(matematicas.PromedioCentrado(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
