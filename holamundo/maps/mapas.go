package main

import "fmt"

func main() {

	monedas := make(map[string]string)
	monedas["mexico"] = "peso mexicano"
	monedas["colombia"] = "peso colombiano"
	monedas["eeuu"] = "eu"

	fmt.Println(monedas)

	delete(monedas, "eeuu")
	fmt.Println(monedas)

	_, existe := monedas["vzla"]

	fmt.Println(existe)

	if existe {
		fmt.Println("si")
	} else {
		fmt.Println("no")
	}

}
