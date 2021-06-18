package main

import (
	"fmt"

	"github.com/CesarDelgadoM/udemyGo/TestBenchmarks/examples/mate"
)

func main() {
	fmt.Println(mate.Sum(2, 3, 4))
	fmt.Println(mate.Sum(4, 7, 8, 9, 123))
}
