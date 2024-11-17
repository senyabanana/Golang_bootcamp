package main

import (
	"log"
	"os"

	"warehouse/internal/adapters"
	"warehouse/internal/repository"
	"warehouse/internal/usecases"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the server port as the first argument")
	}
	serverPort := os.Args[1]

	repo := repository.NewNodeRepository()
	useCase := usecases.NewNodeUseCase(repo)
	server := adapters.NewNodeServer(useCase, serverPort)

	if len(os.Args) == 2 {
		server.SetLeader(true)
		log.Println("This node is the leader")
	}

	server.StartServer()
}
