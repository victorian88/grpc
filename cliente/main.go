package main

import (
	"log"

	pb "github.com/victorian88/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("no se conecto : %v", err)
	}
	defer conn.Close()

	client := pb.NewCodeExecutionServiceClient(conn)
	password := &pb.CodeRequest{
		Pass: "Tecnico88",
	}
	callBireccionalStream(client, password)

}
