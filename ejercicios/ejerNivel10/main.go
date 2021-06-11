package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gen(salir)
	recibir(c, salir)
	fmt.Println("Finalizando...")
}

func gen(salir <-chan int) <-chan int {
	c := make(chan int)
	for i := 0; i < 100; i++ {
		c <- i
	}
	return c
}

func recibir(c <-chan int, salir chan int) {

}
