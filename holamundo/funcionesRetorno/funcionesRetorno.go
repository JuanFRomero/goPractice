package main

import "fmt"

func main() {
	// resultado1 := suma(5, 3)
	// fmt.Println(resultado1)
	// resultado2 := suma(10, 9)
	// fmt.Println(resultado2)

	resultado1 := operacion(9, 5, "suma")
	fmt.Println(resultado1)

	resultado2 := operacion(20, 5, "resta")
	fmt.Println(resultado2)

	resultado3 := operacion(50, 4, "multiplicacion")
	fmt.Println(resultado3)

	resultado4 := operacion(2, 2, "division")
	fmt.Println(resultado4)

	suma, resta, multi := operacion2(3, 9)
	fmt.Println(suma, resta, multi)

	suma, resta, multiplicacion := superOperacion(13, 9)
	fmt.Println(suma, resta, multiplicacion)
}

//operacion (n1,n2,calculo)
// func suma(num1, num2 int) int {
// 	suma := num1 + num2

// 	return suma
// }

func operacion(num1, num2 int, calculo string) int {
	// if calculo == "suma" {
	// 	return num1 + num2
	// } else if calculo == "resta" {
	// 	return num1 - num2
	// } else if calculo == "multiplicacion" {
	// 	return num1 * num2
	// } else {
	// 	return num1 / num2
	// }

	switch {
	case calculo == "suma":
		return num1 + num2
	case calculo == "resta":
		return num1 - num2
	case calculo == "multiplicacion":
		return num1 * num2
	default:
		return 0
	}
}

func operacion2(num1, num2 int) (int, int, int) {
	suma := num1 + num2
	resta := num1 - num2
	multiplicacion := num1 * num2
	return suma, resta, multiplicacion
}

func superOperacion(num1, num2 int) (suma int, resta int, multiplicacion int) {
	suma = num1 + num2
	resta = num1 - num2
	multiplicacion = num1 * num2
	return
}
