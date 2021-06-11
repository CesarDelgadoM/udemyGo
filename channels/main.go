package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
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
	example9()
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

//Context
func example8() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Chequeo de error 1:", ctx.Err())
	fmt.Println("Num de gorutinas 1:", runtime.NumGoroutine())
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Trabajando", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("Chequeo de error 2:", ctx.Err())
	fmt.Println("Num de gorurinas 2:", runtime.NumGoroutine())
	fmt.Println("Cancelando el context...")
	cancel()
	fmt.Println("context cancelado")
	time.Sleep(time.Second * 2)
	fmt.Println("Chequeo de error 3:", ctx.Err())
	fmt.Println("Num de gorutinas 3:", runtime.NumGoroutine())
}

func example9() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		time.Sleep(time.Second * 2)
		if n == 10 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return //deteniendo la rutina
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
