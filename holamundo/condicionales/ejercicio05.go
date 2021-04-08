package main

import "fmt"

func main() {
	//1 2 3 4 5 reprobado
	//6 7 8 9 10 aprobado
	nota := 19

	if nota > 10 {
		fmt.Println("fuera de rango")
	} else {
		if nota > 5 {
			fmt.Println("aprobado")
		} else {
			fmt.Println("desaprobado")
		}
	}

}
