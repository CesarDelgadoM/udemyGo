package main

import (
	"testing"
)

//Tabla de test
func TestMySum(t *testing.T) {
	type Test struct {
		data   []int
		answer int
	}
	test1 := Test{[]int{3, 5, 6}, 14}
	test2 := Test{[]int{3, 3, 3, 3}, 23}
	test3 := Test{[]int{33, 33, 33, 33}, 132}
	test4 := Test{[]int{23, 78, 13, 56}, 156}
	tests := []Test{test1, test2, test3, test4}
	for _, test := range tests {
		r := mySum(test.data...)
		if r != test.answer {
			//Idioma Excepted (esperaba), Got (Obtuvo)
			t.Error("Excepted", test.answer, "Got", r)
		}
	}
}
