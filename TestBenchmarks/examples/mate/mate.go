//Package mate nos ayuda a comprobar que saber sumar
package mate

//Sum suma una cantidad ilimitada de numeros enteros
func Sum(nums ...int) int {
	var sumTotal int
	for _, n := range nums {
		sumTotal = sumTotal + n
	}
	return sumTotal
}
