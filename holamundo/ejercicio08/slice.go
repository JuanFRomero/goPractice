package main

import "fmt"

func main() {
	// var cursos []string

	//agregar informacion al slice

	// cursos = append(cursos, "js")
	// cursos = append(cursos, "c#")
	// cursos = append(cursos, "php")

	cursos := make([]string, 0)

	cursos = []string{"go", "php", "js"}

	fmt.Println(cursos)
}
