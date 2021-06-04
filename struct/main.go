package main

import "fmt"

type persona struct {
	name     string
	lastName string
	age      int
}

type agente struct {
	p   persona
	lpm bool
}

func main() {
	example1()
}

func example1() {
	p1 := persona{
		name:     "cesar",
		lastName: "delgado",
		age:      45,
	}
	p2 := persona{
		name:     "paola",
		lastName: "avella",
		age:      34,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.name, p1.age)
	fmt.Println(p2.name, p2.age)
}

func example2() {
	ag := agente{
		p: persona{
			name:     "Paola",
			lastName: "Avella",
			age:      16,
		},
		lpm: true,
	}
	fmt.Println(ag)
	fmt.Println(ag.p.age)
}

func example3() {
	//struct anonimos
	s1 := struct {
		name     string
		lastName string
		age      int
	}{
		name:     "cesar",
		lastName: "delgado",
		age:      45,
	}
	fmt.Println(s1)
}
