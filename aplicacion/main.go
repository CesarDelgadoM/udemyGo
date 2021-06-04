package main

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	example1()
}

func example1() {
	p1 := persona{
		Nombre:   "Cesar",
		Apellido: "Delgado",
		Edad:     33,
	}
	p2 := persona{
		Nombre:   "Paola",
		Apellido: "Avella",
		Edad:     37,
	}
}
