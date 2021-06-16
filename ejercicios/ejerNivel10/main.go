package main

import (
	"fmt"
	"runtime"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}

func exercise1() {
	salir := make(chan int)
	c := gen(salir)
	recibir(c, salir)
	fmt.Println("Finalizando...")
}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 0
		//close(c)
	}()
	return c
}

func recibir(c <-chan int, salir chan int) {
	for {
		select {
		case <-salir:
			return
		case v := <-c:
			fmt.Println(v)
		}
	}
}

func exercise2() {
	//uso del idioma OK
	c := make(chan int)
	go func() {
		c <- 33
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)
}

func exercise3() {
	c := make(chan int)
	go enviarNumeros(100, &c)
	for v := range c {
		fmt.Println(v)
	}
}

func enviarNumeros(cant int, c *chan int) {
	for i := 0; i < cant; i++ {
		*c <- i
	}
	close(*c)
}

func exercise4() {
	var listChannels []chan int
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println("gorutine:", n)
			cr := cargarChannel(10)
			close(cr)
			listChannels = append(listChannels, cr)
		}(i)
		fmt.Println("Nums gorutines:", runtime.NumGoroutine())
	}
	fmt.Println("Pintamos los canales:", len(listChannels))
	pintarValores(listChannels)
}

func pintarValores(listChannels []chan int) {
	for _, channel := range listChannels {
		for v := range channel {
			fmt.Println("\t", v)
		}
	}
}

func cargarChannel(cant int) chan int {
	c := make(chan int)
	for i := 0; i < cant; i++ {
		c <- i
	}
	return c
}
