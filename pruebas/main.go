package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go exercise7(c)
	}
	for j := 0; j < 20; j++ {
		fmt.Println(<-c)
	}
	wg.Wait()
}

func exercise7(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	wg.Done()
}
