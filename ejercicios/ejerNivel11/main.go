package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

type errorPer struct {
	info string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("Aqui esta el error: %v", ep.info)
}

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := aJSON(p1)
	if err != nil {
		log.Fatalln("Error al realizar marshal", err)
	}
	fmt.Println(string(bs))
}

// aJSON también necesita retorna un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("ERROR al realizar el marshal")
	}
	return bs, err
}

func exercise2() {
	ep := errorPer{
		info: "Necesito dormir mas",
	}
	foo(ep)
}

func foo(err error) {
	fmt.Println("foo corrio - ", err)
}
