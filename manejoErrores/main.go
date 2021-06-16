package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	example1()
	example2()
	example3()
	example4()
	//example5()
	//example6()
	//example7()
	example8()
}

func example1() {
	var resp1, resp2, resp3 string
	fmt.Print("Nombre: ")
	_, err := fmt.Scan(&resp1)
	isErr(err, "Respuesta incorrecta!", true)
	fmt.Print("Comida favorita: ")
	_, err = fmt.Scan(&resp2)
	isErr(err, "Respuesta incorrecta!", false)
	fmt.Print("Deporte favorito: ")
	_, err = fmt.Scan(&resp3)
	isErr(err, "Respuesta incorrecta!", true)
	fmt.Println(resp1, resp2, resp3)
}

func example2() {
	file, err := os.Create("nombres.txt")
	isErr(err, "Error en la creacion del archivo", true)
	defer file.Close() //cerramos el archivo al final de la ejecucion de la funcion
	reader := strings.NewReader("Paola Avella!")
	io.Copy(file, reader)
}

func example3() {
	file, err := os.Open("nombres.txt")
	isErr(err, "Error abriendo el archivo", true)
	defer file.Close()
	bs, err := ioutil.ReadAll(file)
	isErr(err, "Error en la lectura del archivo", true)
	fmt.Println(string(bs))
}

func example4() {
	_, err := os.Open("sin-archivo.txt")
	isErrLog(err, "Error al abrir el achivo -", true)
}

func example5() {
	file, err := os.Create("log.txt")
	isErrLog(err, "Error en la creacion del archivo -", false)
	defer file.Close()
	log.SetOutput(file) //Indicamos el archivo de escritura log
	file2, err := os.Open("sin-archivo.txt")
	isErrLog(err, "Error al abrir el archivo -", false)
	defer file2.Close()
	fmt.Println("Archivo log creado!")
}

func example6() {
	defer foo()
	_, err := os.Open("sin-archivo.txt")
	isErrFatal(err, "Error al abrir el archivo -")
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}

func example7() {
	_, err := os.Open("sin-archivo.txt")
	isErrPanic(err, "Error al abrir el archivo -")
}

func example8() {
	_, err := sqrt(-10)
	isErrFatal(err, "Error en la funcion sqrt().")
	_, err = sqrtFormat(-1)
	isErrFatal(err, "Error en la funcion sqrtFormat().")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("De matematica elemental:" +
			"no hay raiz cuadrada real de un numero negativo")
	}
	return 42, nil
}

func sqrtFormat(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("De matematica elemental:"+
			"no hay raiz cuadrada real de un numero negativo: %v", f)
	}
	return 42, nil
}

func isErr(err error, msg string, exit bool) {
	if err != nil && exit {
		fmt.Println("[ERROR]:", msg, err)
		os.Exit(1)
	} else if err != nil {
		fmt.Println("[ERROR]:", msg, err)
	}
}

func isErrLog(err error, msg string, exit bool) {
	if err != nil && exit {
		log.Println("[ERROR]:", msg, err)
		//os.Exit(1)
	} else if err != nil {
		log.Println("[ERROR]:", msg, err)
	}
}

func isErrFatal(err error, msg string) {
	if err != nil {
		log.Fatalln("[ERROR]:", msg, err)
	}
}

func isErrPanic(err error, msg string) {
	if err != nil {
		log.Panicln("[ERROR]:", msg, err)
	}
}
