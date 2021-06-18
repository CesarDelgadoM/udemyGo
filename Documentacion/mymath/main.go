//Package mymath tiene una funcion sum
package mymath

//Sum suma varios numeros y retorna el valor total
func Sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum = sum + n
	}
	return sum
}
