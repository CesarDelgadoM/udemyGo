package main

import "fmt"

func main() {
	fmt.Println("Suma de 3 4 5 =", mySum(3, 4, 5))
}

func mySum(nums ...int) int {
	var sumTotal int
	for _, n := range nums {
		sumTotal = sumTotal + n
	}
	return sumTotal
}
