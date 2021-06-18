package palabra

import (
	"strings"
)

// No necesitas escribir un ejemplo para esta función
// escribir una prueba para esta es un bono retador; un poco más difícil
func ConteoUso(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

func Conteo(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
