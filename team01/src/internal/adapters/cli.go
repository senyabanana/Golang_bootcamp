package adapters

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"warehouse/internal/usecases"

	"github.com/google/uuid"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	client := usecases.NewClientUsecase()

	fmt.Println("Connected to a database of Warehouse 13")
	fmt.Println("Available commands: GET <key>, SET <key> <value>, DELETE <key>, JOIN <address> <port>")

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()

		if strings.HasPrefix(command, "GET") {
			key := strings.TrimSpace(strings.TrimPrefix(command, "GET"))
			if key == "" {
				fmt.Println("Usage: GET <key>")
				continue
			}

			response, err := client.Get(key)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(response)
			}
		} else if strings.HasPrefix(command, "SET") {
			parts := strings.SplitN(strings.TrimPrefix(command, "SET "), " ", 2)
			if len(parts) != 2 {
				fmt.Println("Usage: SET <key> <value>")
				continue
			}
			key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

			if _, err := uuid.Parse(key); err != nil {
				fmt.Println("Error: Key is not a proper UUID4")
				continue
			}
			err := client.Set(key, value)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Created")
			}
		} else if strings.HasPrefix(command, "DELETE") {
			key := strings.TrimSpace(strings.TrimPrefix(command, "DELETE"))
			if key == "" {
				fmt.Println("Error: key cannot be empty")
				continue
			}

			err := client.Delete(key)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deleted")
			}
		} else if strings.HasPrefix(command, "JOIN") {
			parts := strings.Split(strings.TrimPrefix(command, "JOIN "), " ")
			if len(parts) != 2 {
				fmt.Println("Usage: JOIN <address> <port>")
				continue
			}

			err := client.JoinCluster(parts[0], parts[1])
			if err != nil {
				fmt.Println("Error joining cluster:", err)
			} else {
				fmt.Println("Node joined the cluster successfully")
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
