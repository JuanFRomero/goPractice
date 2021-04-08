package main

import "fmt"

func main() {
	var miPrimeraVariable string         //definiendo variable de manera normal
	miPrimeraVariable = "Aprendiendo Go" //asignando valor a variable

	miSegundaVariable := "Aprendiendo Go otro tipo de variable" //manera corta de crear variable

	miTerceraVariable := 30

	fmt.Println(miPrimeraVariable, miSegundaVariable, miTerceraVariable)

	miCuartaVariable, miQuintaVariable, miSextaVariable := "se pueden declarar variables", "en una misma linea el siguiente es un int", 10

	fmt.Println(miCuartaVariable, miQuintaVariable, miSextaVariable)
}
