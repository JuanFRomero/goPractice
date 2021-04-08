package main

import (
	"fmt"

	"../utils"
)

func main() {
	host := "localhost:8080"
	protocolo := "https"
	uri := "/carritos"
	parametros := map[string]string{"id": "5", "nombre": "galletas"}
	fmt.Println(utils.GenerarUrl(uri, host, protocolo, parametros))
}
