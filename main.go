package main

import (
	"fmt"
	"net/http"
	"productos-api/handlers"
)

func main() {
	http.HandleFunc("/api/productos", handlers.ManejarProductos)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
