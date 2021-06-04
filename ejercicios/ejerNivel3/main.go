package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

func exercise1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	fmt.Println()
}

func exercise3() {
	born := 1996
	yearActual := 2021
	for born < yearActual {
		fmt.Println(born)
		born++
	}
}

func exercise4() {
	born := 1996
	yearActual := 2021
	for {
		if born >= yearActual {
			break
		}
		fmt.Println(born)
		born++
	}
}

func exercise5() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
		i++
	}
}

func exercise6() {
	deporteFav := "baloncesto"
	switch deporteFav {
	case "futbol":
		fmt.Println("ten buenos reflejos")
	case "baloncesto":
		fmt.Println("creceras mucho")
	case "natacion":
		fmt.Println("se rapido como un delfin")
	}
}
