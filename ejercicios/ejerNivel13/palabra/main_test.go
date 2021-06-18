package palabra

import (
	"fmt"
	"testing"
)

func TestConteoUso(t *testing.T) {
	cadena := "Hola mundo en en el lenguaje de de programacion mundo digo go go golang"
	result := map[string]int{
		"en":    2,
		"de":    2,
		"mundo": 2,
		"go":    2,
	}
	m := ConteoUso(cadena)
	fmt.Println(m)
	for key, value := range result {
		for keyM, valM := range m {
			if key == keyM && value != valM {
				t.Error("Excepted:", value, "Got:", valM)
			}
		}
	}
}

func BenchmarkConteoUso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConteoUso("Hola Hola mundo en lenguaje lenguaje Hola en golang golang")
	}
}

func TestConteo(t *testing.T) {
	cadena := "Hola Hola mundo en lenguaje lenguaje Hola en golang golang"
	result := 10
	cont := Conteo(cadena)
	if result != cont {
		t.Error("Expected:", result, "Got:", cont)
	}
}
