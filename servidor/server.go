package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Server struct {
}

func (s *Server) ExecuteCode(param string, reply *string) error {
	// Aquí puedes ejecutar el código en el cliente utilizando el parámetro proporcionado por el servidor
	fmt.Println("Ejecutando código en el cliente:", param)

	// Aquí puedes realizar cualquier procesamiento adicional necesario

	// Establece la respuesta que se enviará de vuelta al servidor
	*reply = "Código ejecutado correctamente en el cliente"

	return nil
}

func main() {
	server := new(Server)

	// Registra el servidor para el servicio de RPC
	rpc.Register(server)

	// Crea un servidor RPC en un puerto específico
	l, err := net.Listen("tcp", ":8095")
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}

	// Acepta conexiones entrantes en el servidor
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error al aceptar la conexión:", err)
		}

		// Inicia el servidor RPC para manejar la conexión entrante
		go rpc.ServeConn(conn)
	}
}
