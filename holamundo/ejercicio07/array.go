package main

import "fmt"

func main() {
	var cursos [3]string

	cursos[0] = "go"
	cursos[1] = "javascript"
	cursos[2] = "python"

	fmt.Println(cursos[1])
	fmt.Println(len(cursos))

	//todos elementos en conjunto
	curso := [3]string{"go", "php", "c#"}
	// fmt.Println(curso)

	//recorrer con for

	for i := 0; i < len(curso); i++ {
		fmt.Println(curso[i])
	}

	for i, valor := range curso {
		fmt.Println(i, valor)
	}
}
