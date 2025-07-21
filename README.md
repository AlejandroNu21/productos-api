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
Abre la terminal y ejecuta tu programa:
- go run main.go
- Servidor corriendo en http://localhost:8080

3. **Abrir Postman y crear nueva solicitud**:
Crear un producto (POST):
- En Postman, haz clic en "New" > HTTP Request
- Método: POST
- URL: http://localhost:8080/api/productos
- Ve a la pestaña "Body", selecciona "raw" y elige tipo JSON.

Escribir algo como:
{
  "nombre": "Mouse",
  "descripcion": "Gamer",
  "precio": 25,
  "stock": 10
}


**Ver todos los productos (GET)**
- Método: GET
- URL: http://localhost:8080/api/productos?id=1 (cambiar el Id según el producto)
- Haz clic en "Send"


**Actualizar un producto (PUT)**
- Método: PUT
- URL: http://localhost:8080/api/productos?id=1
- Ve a Body > raw > JSON y escribe algo como:

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
- URL: http://localhost:8080/api/productos?id=1
- Haz clic en "Send"
- saldrá el mensaje: "Producto eliminado"
