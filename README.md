# Productos API 

Esta es una API RESTful sencilla desarrollada en **Go (Golang)** que simula un backend para una tienda en línea de productos tecnológicos. Permite realizar operaciones CRUD (crear, leer, actualizar y eliminar) sobre productos del inventario.

---

## Características

- Consultar todos los productos
- Consultar un producto por ID
- Agregar un nuevo producto
- Editar un producto existente
- Eliminar un producto del inventario

---

## Requisitos

- [Go](https://go.dev/dl/) 1.18 o superior
- Visual Studio Code (opcional)
- Instalar Postman o cualquier herramienta para probar APIs

---

## Cómo ejecutar el proyecto

1. **Clona el repositorio**:
   ```bash
   git clone https://github.com/tu-usuario/productos-api.git
   cd productos-api

2. **Abrir Archivo en Visual Studio Code**:
Abre la terminal y ejecuta la API:
- go run main.go
- Servidor corriendo en http://localhost:8080

3. **Abrir Postman y crear nueva solicitud**:
Crear un producto (POST):
- En Postman, haz clic en "New" > HTTP Request
- Método: POST
- URL: http://localhost:8080/api/productos
- Ir a la pestaña "Body", seleccionar "raw" y elegir tipo JSON.

Escribir algo como:
{
  "nombre": "Mouse",
  "descripcion": "Gamer",
  "precio": 25,
  "stock": 10
}


**Ver todos los productos (GET)**
- Método: GET
- URL: http://localhost:8080/api/productos
- Hacer clic en "Send"


**Actualizar un producto (PUT)**
- Método: PUT
- URL: http://localhost:8080/api/productos?id=1 (Cambiar id por el que se desea actualizar)
- Ir a Body > raw > JSON y escribir algo como:

{
  "nombre": "Mouse Pro XL",
  "descripcion": "Gamer",
  "precio": 30,
  "stock": 10
}

- Hacer clic en "Send"
- saldrá el mensaje: "Producto actualizado"


**Eliminar un producto (DELETE)**
- Método: DELETE
- URL: http://localhost:8080/api/productos?id=1 (Cambiar id por el que se desea eliminar)
- Hacer clic en "Send"
- saldrá el mensaje: "Producto eliminado"
