package main

import (
	"log"
	"net"

	srv "all-together/internal/server"
	pb "all-together/internal/transmitter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":50051"

func main() {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTransmitterServiceServer(grpcServer, &srv.Server{})
	log.Println("server is running on port :50051")

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
