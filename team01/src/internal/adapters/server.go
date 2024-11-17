package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"warehouse/internal/usecases"

	"github.com/google/uuid"
)

type NodeServer struct {
	useCase    *usecases.NodeUseCase
	serverPort string
	isLeader   bool
	mu         sync.Mutex
}

func NewNodeServer(useCase *usecases.NodeUseCase, serverPort string) *NodeServer {
	return &NodeServer{
		useCase:    useCase,
		serverPort: serverPort,
	}
}

func (c *NodeServer) StartServer() {
	http.HandleFunc("/get", c.handleGet)
	http.HandleFunc("/set", c.handleSet)
	http.HandleFunc("/delete", c.handleDelete)
	http.HandleFunc("/heartbeat", c.handleHeartbeat)
	http.HandleFunc("/join", c.handleJoinCluster)

	go c.useCase.MonitorNodes()

	log.Printf("Server starting on port %s", c.serverPort)
	log.Fatal(http.ListenAndServe(":"+c.serverPort, nil))
}

func (c *NodeServer) SetLeader(isLeader bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.isLeader = isLeader
}

func (c *NodeServer) handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Error: key cannot be empty", http.StatusBadRequest)
		return
	}

	value, err := c.useCase.GetData(key)
	if err != nil {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, value)
}

func (c *NodeServer) handleSet(w http.ResponseWriter, r *http.Request) {
	if !c.isLeader {
		http.Error(w, "Current node is not the leader", http.StatusForbidden)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Error: key cannot be empty", http.StatusBadRequest)
		return
	}

	if _, err := uuid.Parse(key); err != nil {
		http.Error(w, "Error: Key is not a proper UUID4", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	value := string(body)

	err = c.useCase.SetData(key, value, c.serverPort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Created/Updated (replicated)")
}

func (c *NodeServer) handleDelete(w http.ResponseWriter, r *http.Request) {
	if !c.isLeader {
		http.Error(w, "Current node is not the leader", http.StatusForbidden)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Error: key cannot be empty", http.StatusBadRequest)
		return
	}

	err := c.useCase.DeleteData(key, c.serverPort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted (replicated)")
}

func (c *NodeServer) handleHeartbeat(w http.ResponseWriter, r *http.Request) {
	response := c.useCase.GetHeartbeatStatus()
	json.NewEncoder(w).Encode(response)
}

func (c *NodeServer) handleJoinCluster(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	port := r.URL.Query().Get("port")
	if address == "" || port == "" {
		http.Error(w, "Error: address and port cannot be empty", http.StatusBadRequest)
		return
	}

	err := c.useCase.JoinCluster(address, port)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Node %s:%s joined the cluster", address, port)
}
