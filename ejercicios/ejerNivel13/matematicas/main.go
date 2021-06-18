package matematicas

import "sort"

// PromedioCentrado computa el promedio de una lista de números
// después de eliminar el valor más grande y el más paqueño.
func PromedioCentrado(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
