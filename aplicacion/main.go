package main

import (
	"encoding/json"
	"fmt"
	//"golang.org/x/crypto/bcrypt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	example1()
	example2()
	//example3()
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
	personas := []persona{p1, p2}
	fmt.Println(personas)
	bs, err := json.Marshal(personas)
	isErr(err, "Error en el Marshal:")
	fmt.Println(string(bs))
}

func example2() {
	js := `[{"Nombre":"Cesar","Apellido":"Delgado","Edad":33},{"Nombre":"Paola","Apellido":"Avella","Edad":37}]`
	bs := []byte(js)
	fmt.Printf("%T\n", js)
	fmt.Printf("%T\n", bs)
	var personas []persona
	err := json.Unmarshal(bs, &personas)
	isErr(err, "Error en el unmarshal:")
	fmt.Println("Toda la data", personas)
	for _, p := range personas {
		fmt.Println(p)
	}
}

/*
func example3() {
	pass := `pass1234`
	bytepass, err := bcrypt.GenerateFromPassWord([]byte(pass), 4)
	isErr(err, "Error en la encriptacion de la password")
	fmt.Println(bytepass)
}
*/

func isErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
