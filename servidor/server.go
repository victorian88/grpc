package main

import (
	//"context"
	"io"
	"log"

	//"google.golang.org/grpc"

	pb "github.com/victorian88/grpc" // Reemplaza con la ruta de importación correcta de tu archivo .pb.go generado
)

/*type helloServer struct {
	pb.CodeExecutionServiceServer
}*/

func (s *helloServer) Bidireccional(stream pb.CodeExecutionService_BidireccionalServer) error {

	for { // Enviar el primer mensaje al cliente
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("se octuvo respuesta clave: %v", req.Pass)
		res := &pb.CodeResponse{
			Message: "Hello" + req.Pass,
		}
		if err := stream.Send(res); err != nil {

			return err
		}
	}
}
