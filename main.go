package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
}

var productos = []Producto{}

func main() {
	http.HandleFunc("/api/productos", manejarProductos)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func manejarProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Query().Get("id")

	switch r.Method {
	case "GET":
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			for _, p := range productos {
				if p.ID == id {
					json.NewEncoder(w).Encode(p)
					return
				}
			}
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(productos)

	case "POST":
		var nuevo Producto
		if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
			http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
			return
		}
		nuevo.ID = len(productos) + 1
		productos = append(productos, nuevo)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nuevo)

	case "PUT":
		if idParam == "" {
			http.Error(w, "Se requiere el parámetro 'id'", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var actualizado Producto
		if err := json.NewDecoder(r.Body).Decode(&actualizado); err != nil {
			http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
			return
		}

		encontrado := false
		for i, p := range productos {
			if p.ID == id {
				actualizado.ID = id
				productos[i] = actualizado
				encontrado = true
				break
			}
		}

		if !encontrado {
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Producto actualizado"})

	case "DELETE":
		if idParam == "" {
			http.Error(w, "Se requiere el parámetro 'id'", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		encontrado := false
		for i, p := range productos {
			if p.ID == id {
				productos = append(productos[:i], productos[i+1:]...)
				encontrado = true
				break
			}
		}

		if !encontrado {
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Producto eliminado"})

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
