package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// Establece la conexión con el cliente (que ahora actúa como servidor)
	conn, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error al establecer la conexión con el cliente:", err)
	}

	// Define el parámetro que se enviará al cliente
	param := "parámetro del servidor"

	// Variable para almacenar la respuesta del cliente
	var reply string

	// Llama a la función "ExecuteCode" en el cliente (que ahora actúa como servidor) y pasa el parámetro
	err = conn.Call("Server.ExecuteCode", param, &reply)
	if err != nil {
		log.Fatal("Error al llamar a la función en el cliente:", err)
	}

	// Muestra la respuesta del cliente
	fmt.Println("Respuesta del cliente:", reply)
}
