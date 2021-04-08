package main

import (
	"fmt"
)

func main() {
	host := "localhost:8080"
	protocolo := "https"
	uri := "/carritos"
	parametros := map[string]string{"id": "5", "nombre": "galletas"}
	fmt.Println(generarUrl(uri, host, protocolo, parametros))

	// u, e := url.Parse("/carritos")

	// if e != nil {
	// 	panic("ocurrio error")
	// }
	// u.Host = host
	// u.Scheme = protocolo

	// mapa := u.Query()
	// mapa.Add("id", "5")
	// mapa.Add("nombre", "gaseosa")
	// u.RawQuery = mapa.Encode()

	// fmt.Println(u.String())

}

// func generarUrl(uri, host, protocolo string, mapa map[string]string) string {
// 	u, _ := url.Parse(uri)
// 	u.Host = host
// 	u.Scheme = protocolo
// 	mapaFuncion := u.Query()

// 	for key, value := range mapa {
// 		mapaFuncion.Add(key, value)
// 	}
// 	u.RawQuery = mapaFuncion.Encode()
// 	return u.String()
// }
