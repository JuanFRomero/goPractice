package main

import (
	"errors"
	"fmt"
)

func main() {
	respuesta, err := operacion(5, -8, "multiplicacion")
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(respuesta)
	}
}

func operacion(num1, num2 int, calculo string) (resultado int, err error) {

	if num1 < 0 {
		err = errors.New("mi primero numero es 0")
	} else if num2 < 0 {
		err = errors.New("numero 2 es negativo")
	} else {
		switch {
		case calculo == "suma":
			resultado = num1 + num2
		case calculo == "resta":
			resultado = num1 - num2
		case calculo == "multiplicacion":
			resultado = num1 * num2
		default:
			resultado = 0
		}
	}
	return
}
