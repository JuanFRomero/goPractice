package main

import "fmt"

func main() {

	// func() {
	// 	fmt.Println("funcion anonima")
	// }() funcion anonima autoejecutable

	// saludo := func() {
	// 	fmt.Println("otra funcion anonima")
	// }
	// saludo()
	// nombres(1, "go", "js", "c#")

	a := 4
	triplicarValor(&a)
	fmt.Println(a)
}

// func nombres(numero int, cursos ...string) {
// 	for _, c := range cursos {
// 		fmt.Println("mi curso es", numero, c)
// 	}
// }

func triplicarValor(n1 *int) {
	*n1 = *n1 * 3
	fmt.Println("el triple valor", *n1)
}
