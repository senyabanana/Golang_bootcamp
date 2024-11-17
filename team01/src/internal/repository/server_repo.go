package repository

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"warehouse/internal/entity"
)

type ServerRepository interface {
	GetData(key string) (string, bool)
	SetData(key, value string)
	DeleteData(key string)
	ReplicateToNodes(key, value, serverPort string, replicationFactor int)
	JoinCluster(address, port string) error
	GetHeartbeatStatus() map[string]interface{}
	MonitorNodes()
}

type NodeRepository struct {
	dataStore map[string]string
	nodes     map[string]*entity.Node
	mu        sync.Mutex
}

func NewNodeRepository() *NodeRepository {
	return &NodeRepository{
		dataStore: make(map[string]string),
		nodes:     make(map[string]*entity.Node),
	}
}

func (r *NodeRepository) GetData(key string) (string, bool) {
	value, exists := r.dataStore[key]
	return value, exists
}

func (r *NodeRepository) SetData(key, value string) {
	r.dataStore[key] = value
}

func (r *NodeRepository) DeleteData(key string) {
	delete(r.dataStore, key)
}

func (r *NodeRepository) ReplicateToNodes(key, value, serverPort string, replicationFactor int) {
	replicatedCount := 0
	for port := range r.nodes {
		if port == serverPort || replicatedCount >= replicationFactor {
			continue
		}
		url := fmt.Sprintf("http://127.0.0.1:%s/set?key=%s", port, key)
		log.Printf("Replicating to node %s", port)

		req, err := http.NewRequest("POST", url, strings.NewReader(value))
		if err != nil {
			log.Printf("Error creating request for node %s: %v", port, err)
			continue
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Printf("Failed to replicate to node %s: %v", port, err)
		} else {
			replicatedCount++
		}
		defer resp.Body.Close()
	}
}

func (r *NodeRepository) JoinCluster(address, port string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.nodes[port]; !exists {
		r.nodes[port] = &entity.Node{
			Address:       address,
			Port:          port,
			IsLeader:      false,
			LastHeartbeat: time.Now(),
		}
		log.Printf("Node %s:%s joined the cluster", address, port)
	}
	return nil
}

func (r *NodeRepository) GetHeartbeatStatus() map[string]interface{} {
	status := make(map[string]interface{})
	status["nodes"] = r.nodes
	return status
}

//func (r *NodeRepository) MonitorNodes() {
//	for port, node := range r.nodes {
//		if time.Since(node.LastHeartbeat) > 10*time.Second {
//			log.Printf("Node %s is down", port)
//			r.mu.Lock()
//			delete(r.nodes, port)
//			r.mu.Unlock()
//		}
//	}
//}

func (r *NodeRepository) MonitorNodes() {
	for port, node := range r.nodes {
		if time.Since(node.LastHeartbeat) > 10*time.Second {
			log.Printf("Node %s is down", port)
			r.mu.Lock()
			delete(r.nodes, port)
			r.mu.Unlock()
			// Логика выбора нового лидера
			if node.IsLeader {
				log.Println("Leader is down, electing a new leader")
				for _, n := range r.nodes {
					n.IsLeader = true
					log.Printf("New leader elected: %s", n.Port)
					break
				}
			}
		}
	}
}
