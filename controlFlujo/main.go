package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
}

func example1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("El ciclo externo %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tEl ciclo interno %d\n", j)
		}
	}
}

func example2() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	i = 0
	for {
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func example3() {
	n := 0
	for {
		n++
		if n > 101 {
			break
		} else if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}
}

func example4() {
	//Table ASCII
	for i := 32; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\n", i, i, i)
	}
}

func example5() {
	switch {
	case 2 == 4:
		fmt.Println("No imprime")
		fallthrough //vuelve verdadero el siguiente caso asi no se cumpla
	case 3 == 3:
		fmt.Println("Si imprime")
	case 4 == 5, 8 == 8:
		fmt.Println("Imprime la primera condicion que cumpla")
	default:
		fmt.Println("Imprime desde default")
	}
}
