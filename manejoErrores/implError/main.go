package main

import (
	"fmt"
	"log"
)

type ubicacionError struct {
	lat  string
	long string
	err  error
}

func (e ubicacionError) Error() string {
	return fmt.Sprintf("Un error de concepto matematico ha "+
		"ocurrido en %v %v %v", e.lat, e.long, e.err)
}

func main() {
	_, err := sqrt(-1)
	isErr(err)
}

func isErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Error mat: raiz cuadrada de "+
			"numero negativo: %v", f)
		return 0, ubicacionError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  nme,
		}
	}
	return 42, nil
}
