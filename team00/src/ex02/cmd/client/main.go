package main

import (
	"flag"
	"log"

	"report/internal/client"
	"report/internal/database"
	pb "report/internal/transmitter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcPort = ":50051"

func main() {
	var k float64

	flag.Float64Var(&k, "k", 2.0, "STD anomaly coefficient")
	flag.Parse()

	db := database.ConnectDB()

	conn, err := grpc.NewClient(grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewTransmitterServiceClient(conn)

	client.AnomalyDetector(grpcClient, db, k)
}
