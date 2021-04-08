package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := 2
	num2 := 2

	//suma

	suma := num1 + num2

	resta := num1 - num2

	multiplicacion := num1 * num2

	division := num1 / num2

	fmt.Printf("suma %d \n, resta %d \n, multipli %d \n, div %d \n", suma, resta, multiplicacion, division)

	//convertir de cadena a entero y de entero a cadena

	num3 := "15"

	numcon1, _ := strconv.Atoi(num3)

	sumaEspecial := numcon1 + num1

	fmt.Println(sumaEspecial)

}
