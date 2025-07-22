package handlers

import (
	"encoding/json"
	"net/http"
	"productos-api/data"
	"productos-api/models"
	"strconv"
)

func ManejarProductos(w http.ResponseWriter, r *http.Request) {
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
			for _, p := range data.Productos {
				if p.ID == id {
					json.NewEncoder(w).Encode(p)
					return
				}
			}
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(data.Productos)

	case "POST":
		var nuevo models.Producto
		if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
			http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
			return
		}
		nuevo.ID = len(data.Productos) + 1
		data.Productos = append(data.Productos, nuevo)
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

		var actualizado models.Producto
		if err := json.NewDecoder(r.Body).Decode(&actualizado); err != nil {
			http.Error(w, "Error en el formato del JSON", http.StatusBadRequest)
			return
		}

		encontrado := false
		for i, p := range data.Productos {
			if p.ID == id {
				actualizado.ID = id
				data.Productos[i] = actualizado
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
		for i, p := range data.Productos {
			if p.ID == id {
				data.Productos = append(data.Productos[:i], data.Productos[i+1:]...)
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
