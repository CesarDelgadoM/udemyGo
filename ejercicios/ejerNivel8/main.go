package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
}

type DichosPersona struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	p1 := Persona{
		Nombre: "Paola",
		Edad:   27,
	}
	p2 := Persona{
		Nombre: "Cesar",
		Edad:   25,
	}
	personas := []Persona{p1, p2}
	bs, err := json.Marshal(personas)
	isError(err, "Error en Marshal")
	fmt.Println(string(bs))
}

func exercise2() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`
	bs := []byte(s)
	var dichosPer []DichosPersona
	err := json.Unmarshal(bs, &dichosPer)
	isError(err, "Error en el Unmarshal")
	fmt.Println(dichosPer)
}

func isError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
