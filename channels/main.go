package main

import (
	"fmt"
	"sync"
)

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
	example8()
}

func example1() {
	//Creacion del channel sin buffer
	ch := make(chan int)
	go enviarDato(ch)
	fmt.Println(<-ch) //Recibimos el dato que envia en channel
}

func enviarDato(ch chan int) {
	ch <- 33
}

func example2() {
	//Creacion de un channel con buffer
	ch := make(chan int, 2)
	ch <- 33
	ch <- 34
	fmt.Println(<-ch)
	fmt.Println(<-ch) //Accedemos al otro dato que quedo en el channel
}

func example3() {
	//Channel bidireccional
	chb := make(chan int)
	//Channel recive
	chr := make(<-chan int)
	//Channel send
	chs := make(chan<- int)

	fmt.Printf("chb\t%T\n", chb)
	fmt.Printf("chr\t%T\n", chr)
	fmt.Printf("chs\t%T\n", chs)
}

func example4() {
	chb := make(chan int)

	go send(chb)
	recive(chb)

	fmt.Println("Finalizando...")
}

func send(chs chan<- int) {
	chs <- 67
}

func recive(chs <-chan int) {
	fmt.Println(<-chs)
}

func example5() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

var wg sync.WaitGroup

func example6() {
	wg.Add(2)
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)
	go enviar(par, impar, salir)
	go recibir(par, impar, salir)
	wg.Wait()
	fmt.Println("Finalizando...")
}

func enviar(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	s <- 0
	wg.Done()
}

func recibir(p, i, s <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("Canal par", v)
		case v := <-i:
			fmt.Println("Canal impar", v)
		case v := <-s:
			fmt.Println("Canal salir", v)
			wg.Done()
			return
		}
	}
}

//idioma coma OK
func example7() {
	c := make(chan int)
	go func() {
		c <- 33
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}

//Rob Pike
func example8() {
	c := fanIn(boring("Paola"), boring("Cesar"))
	for i := 0; i < 10; i++ {
		fanIn(boring("Paola"), boring("Cesar"))
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		c <- <-input1
	}()
	go func() {
		c <- <-input2
	}()
	return c
}
