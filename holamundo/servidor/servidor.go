package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sirve ruta")
	})
	http.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pagina productos")
		fmt.Println(r.URL)
		mapa := r.URL.Query()
		mapa.Add("precio", "10")
		fmt.Println(mapa.Encode())
		fmt.Println(r.URL.RawQuery)
		fmt.Println("id", r.URL.Query().Get("idProducto"))
		fmt.Println("nombre", r.URL.Query().Get("nombre"))
	})
	http.HandleFunc("/categorias", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "pagina categorias")
		http.Redirect(w, r, "/", 300)
	})
	http.HandleFunc("/nada", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "pagina categorias")
		http.NotFound(w, r)
	})
	http.ListenAndServe(":7000", nil)
}
