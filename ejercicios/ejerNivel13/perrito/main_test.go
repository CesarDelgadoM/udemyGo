package perrito

import (
	"fmt"
	"math/rand"
	"testing"
)

type Test struct {
	year   int
	result int
}

func TestEdad(t *testing.T) {
	test1 := Test{year: 4, result: 28}
	test2 := Test{year: 8, result: 56}
	tests := []Test{test1, test2}
	for _, test := range tests {
		r := Edad(test.year)
		if r != test.result {
			t.Error("Excepted:", test.result, "Got:", r)
		}
	}
}

func TestEdadDos(t *testing.T) {
	test1 := Test{year: 2, result: 14}
	test2 := Test{year: 4, result: 28}
	tests := []Test{test1, test2}
	for _, test := range tests {
		r := EdadDos(test.year)
		if r != test.result {
			t.Error("Excepted:", test.result, "Got:", r)
		}
	}
}

func ExampleEdad() {
	fmt.Println(Edad(34))
	//Output:
	//238
}

func ExampleEdadDos() {
	fmt.Println(EdadDos(123))
	//Output:
	//861
}

func BenchmarkEdad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		year := rand.Intn(100)
		fmt.Println(year)
		Edad(year)
	}
}

func BenchmarkEdadDos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		year := rand.Intn(100)
		fmt.Println(year)
		EdadDos(year)
	}
}
