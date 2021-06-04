package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
	example8()
	example9()
	example10()
	example11()
}

func example1() {
	var x [5]int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
}

func example2() {
	//Literal compuesto (composite literal)
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Printf("%T", x)
}

func example3() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x[1:4])
}

func example4() {
	x := []int{1, 2, 3}
	y := []int{34, 67, 12, 78}
	fmt.Println(x)
	x = append(x, 3, 4, 5)
	fmt.Println(x)
	x = append(x, y...) //agreamos los elementos del slice 'y'
	fmt.Println(x)
}

func example5() {
	x := []int{4, 5, 7, 8, 42, 44, 55, 66, 333, 444, 666}
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func example6() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[5] = 100
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 200)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func example7() {
	cd := []string{"Cesar", "Delgado", "Basketball", "Crossfit", "Calestenia"}
	fmt.Println(cd)
	pa := []string{"Paola", "Avella", "Bailar", "Musica", "Vino"}
	fmt.Println(pa)
	vp := [][]string{cd, pa}
	fmt.Println(vp)
}

func example8() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	y := append(x[:2], x[3:]...) //Se utiliza el mismo arreglo subyacente
	fmt.Println(x)
	fmt.Println(y)
}

func example9() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)
	fmt.Println(m["Batman"])
	fmt.Println(m["Robin"])

	v, ok := m["Cesar"]
	fmt.Println(v)
	fmt.Println(ok)

	if _, ok := m["Paola"]; !ok {
		fmt.Println("Paola no existe en el mapa")
	}
}

func example10() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	m["Paola"] = 27
	fmt.Println(m)
	for key, value := range m {
		fmt.Println(key, ",", value)
	}
}

func example11() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)
	if _, ok := m["Batman"]; ok {
		fmt.Println("Borrando elemento")
		delete(m, "Batman")
	}
	fmt.Println(m)
}
