package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/victorian88/grpc"
)

func callBireccionalStream(client pb.CodeExecutionServiceClient, password *pb.CodeRequest) {
	log.Printf(" conexion bidireccional empezo")
	stream, err := client.Bidireccional(context.Background())
	if err != nil {
		log.Fatalf("No se pudo enviar la contraseña %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {

				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	req := &pb.CodeRequest{
		Pass: password.Pass,
	}

	if err := stream.Send(req); err != nil {
		log.Fatalf("error while sending %v", err)
	}
	time.Sleep(2 * time.Second)
	stream.CloseSend()
	<-waitc
	log.Println("Se terminó la conexion")

}
