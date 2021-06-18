package main

import (
	"fmt"

	"github.com/CesarDelgadoM/udemyGo/ejercicios/ejerNivel12/perro"
)

func main() {
	yearsHuman := 12
	yearsDog := perro.GetYearsDog(yearsHuman)
	fmt.Println(yearsDog)
}
