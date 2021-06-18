// Package perrito nos permite entender mejor los perros.
package perrito

// Edad convierte años humanos en años perros.
func Edad(n int) int {
	return n * 7
}

// EdadDos convierte años humanos en años perros.
func EdadDos(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
