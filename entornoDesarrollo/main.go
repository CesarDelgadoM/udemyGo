package main

import (
	"fmt"
	"runtime"
)

func main() {
	example1()
}

//Compilacion cruzada
func example1() {
	fmt.Printf("Sistema operativo %v\nArquitectura: %v\n", runtime.GOOS, runtime.GOARCH)
	//Para compilar indicando el OS y la arquitectura se lanza el siguiente comando:
	//GOOS=linux GOARCH=amd64 go build main.go
}
