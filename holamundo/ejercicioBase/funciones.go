package main

import "fmt"

func main() {
	saludas("de juan", "romero", "guerrero", 1)
}

func saludas(nombre, apellido, apellidoM string, numero int) {
	fmt.Println("hola funcion numero", numero, nombre, apellido, apellidoM)
}
