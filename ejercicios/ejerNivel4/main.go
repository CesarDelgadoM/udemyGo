package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
}

func exercise1() {
	arr := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", arr)
}

func exercise2() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", arr)
}

func exercise3() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])
}

func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func exercise6() {
	x := make([]string, 3, 5)
	x = []string{"Paipa", "Duitama", "Tunja"}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i+1, x[i])
	}
}

func exercise7() {
	p := [][]string{{"Batman", "Jefe", "Vestido de negro"},
		{"Robin", "Ayudame", "Vestido de colores"}}
	for _, v := range p {
		fmt.Println(v)
	}
	for _, v := range p {
		for _, vi := range v {
			fmt.Println(vi)
		}
	}
}

func exercise8() {
	m := map[string][]string{
		"cesar_delgado": {"computadoras", "videojuegos", "maquinas"},
		"paola_avella":  {"Universo", "vintage", "peliculas"},
	}
	for i, v := range m {
		fmt.Println(i)
		for i, j := range v {
			fmt.Printf("\t%d %s\n", i, j)
		}
	}
}

func exercise9() {
	m := map[string][]string{
		"cesar_delgado": {"computadoras", "videojuegos", "maquinas"},
		"paola_avella":  {"Universo", "vintage", "peliculas"},
	}
	if _, ok := m["cesar_delgado"]; ok {
		delete(m, "cesar_delgado")
	}
	fmt.Println(m)
}
