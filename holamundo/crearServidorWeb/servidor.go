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
