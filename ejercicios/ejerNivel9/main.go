package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Persona struct {
}

func (p *Persona) Hablar() {
	fmt.Println("Estoy hablando")
}

type Humano interface {
	Hablar()
}

func diAlgo(h Humano) {
	h.Hablar()
}

var wg sync.WaitGroup

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	wg.Add(2)
	go saludar()
	go despedir()
	wg.Wait()
}

func exercise2() {
	p := Persona{}
	fmt.Printf("PUEDES pasar un valor de tipo %T a diAlgo: ", p)
	diAlgo(&p)
	fmt.Println("NO puedes pasar un valor de tipo persona a diAlgo")
	p.Hablar()
}

func exercise3() {
	cont := 0
	const gr = 50
	var wg sync.WaitGroup
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			v := cont
			runtime.Gosched()
			v++
			cont = v
			wg.Done()
		}()
		fmt.Println("Numero de goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Valor final del contador:", cont)
}

func saludar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Hola")
	}
	wg.Done()
}

func despedir() {
	for i := 0; i < 10; i++ {
		fmt.Println("Adios")
	}
	wg.Done()
}
