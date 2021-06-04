package main

import "fmt"

type persona struct {
	name        string
	last        string
	iceCreamFav []string
}

type vehicle struct {
	numDoors int
	color    string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	p1 := persona{
		name:        "Paola",
		last:        "Avella",
		iceCreamFav: []string{"Frutos rojos", "Chocolate"},
	}
	p2 := persona{
		name:        "Cesar",
		last:        "Delgado",
		iceCreamFav: []string{"lulo", "brownie"},
	}
	fmt.Println(p1)
	fmt.Println("Helados favoritos de", p1.name)
	for _, hf := range p1.iceCreamFav {
		fmt.Println("\t", hf)
	}
	fmt.Println(p2)
	fmt.Println("Helados favoritos de", p2.name)
	for _, hf := range p2.iceCreamFav {
		fmt.Println("\t", hf)
	}
}

func exercise2() {
	p1 := persona{
		name:        "Paola",
		last:        "Avella",
		iceCreamFav: []string{"Frutos rojos", "Chocolate"},
	}
	p2 := persona{
		name:        "Cesar",
		last:        "Delgado",
		iceCreamFav: []string{"lulo", "brownie"},
	}
	m1 := map[string]persona{
		p1.last: p1,
		p2.last: p2,
	}
	for _, p := range m1 {
		fmt.Println(p)
		fmt.Println("Helados favorios de", p.name)
		for _, hf := range p.iceCreamFav {
			fmt.Println("\t", hf)
		}
	}
}

func exercise3() {
	truck1 := truck{
		vehicle: vehicle{
			numDoors: 2,
			color:    "vinotinto",
		},
		fourWheels: false,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			numDoors: 4,
			color:    "green",
		},
		luxury: true,
	}
	fmt.Println("numero de puertas del camion", truck1.numDoors)
	fmt.Println("color del camion", truck1.color)
	fmt.Println("tiene cuatro ruedas?", truck1.fourWheels)
	fmt.Println("numero de puertas del sedan", sedan1.numDoors)
	fmt.Println("color del sedan", sedan1.color)
	fmt.Println("es lujoso?", sedan1.luxury)
}
