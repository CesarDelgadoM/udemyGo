package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go add(c1)
	go fanOutIn(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func add(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		fmt.Println(<-c1)
		wg.Add(1)
		go func(v int) {
			c2 <- trabajoConsumeTiempo(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func trabajoConsumeTiempo(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
