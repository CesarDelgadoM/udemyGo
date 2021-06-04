package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

func exercise1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

var x int
var y string
var z bool

func exercise2() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	//Se denominan valores cero
}

func exercise3() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)
}

type temperature int

var t temperature

func exercise4() {
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	t = 33
	fmt.Println(t)
}

var p int

func exercise5() {
	p = int(t)
	fmt.Println(p)
}
