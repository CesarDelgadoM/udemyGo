package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	exercise1()
	wg.Add(1) //agregamos un goroutine
	exercise2()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() //esperamos las goroutine que finalicen
}

func exercise1() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func exercise2() {
	go foo()
	bar()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
