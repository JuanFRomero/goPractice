package main

import "fmt"

//almacenar cursos(id,nombre,apellido)
type Curso struct {
	id       int
	nombre   string
	apellido string
}

func main() {
	// var ocurso Curso

	// ocurso.id = 1
	// ocurso.nombre = "Juan"
	// ocurso.apellido = "Romero"
	// fmt.Println(ocurso)
	// fmt.Println(ocurso.nombre)
	// fmt.Println(ocurso.id)
	// fmt.Println(ocurso.apellido)

	ocurso := Curso{id: 1, nombre: "juan", apellido: "romero"}
	// ocurso.id = 1
	// ocurso.nombre = "Juan"
	// ocurso.apellido = "Romero"
	fmt.Println(ocurso)
}
