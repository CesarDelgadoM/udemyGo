package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

func exercise1() {
	num := 33
	fmt.Printf("%d\n%b\n%x\n", num, num, num)
}

func exercise2() {
	a := (33 == 33)
	b := (56 <= 12)
	c := (23 >= 78)
	d := (34 < 90)
	e := (12 > 4)
	f := (56 != 57)
	println(a, b, c, d, e, f)
}

const x = 30
const y int = 33

func exercise3() {
	fmt.Println(x)
	fmt.Println(y)
}

func exercise4() {
	x := 45
	fmt.Printf("%d %b %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d %b %#x\n", y, y, y)
}

func exercise5() {
	s := `Esto es un string literal
	no interpretado "¿lo ves?"`
	fmt.Println(s)
}

const (
	yearActual = 2021
	primerYear = iota + yearActual
	segYear
	terYear
	cuartoYear
)

func exercise6() {
	fmt.Println(primerYear, segYear,
		terYear, cuartoYear)
}
