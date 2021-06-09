package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	exercise1()
}

func exercise1() {
	wg.Add(2)
	go saludar()
	go despedir()
	wg.Wait()
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
