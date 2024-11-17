package usecases

import (
	"fmt"
	"log"
	"sync"
	"time"

	"warehouse/config"
	"warehouse/internal/repository"
)

type NodeUseCase struct {
	repo repository.ServerRepository
	mu   sync.Mutex
}

func NewNodeUseCase(repo repository.ServerRepository) *NodeUseCase {
	return &NodeUseCase{
		repo: repo,
	}
}

func (u *NodeUseCase) GetData(key string) (string, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	value, exists := u.repo.GetData(key)
	if !exists {
		return "", fmt.Errorf("key not found")
	}
	return value, nil
}

func (u *NodeUseCase) SetData(key, value, serverPort string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.repo.SetData(key, value)
	u.repo.ReplicateToNodes(key, value, serverPort, config.ReplicationFactor)
	log.Printf("Data set and replicated for key %s", key)
	return nil
}

func (u *NodeUseCase) DeleteData(key, serverPort string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.repo.DeleteData(key)
	u.repo.ReplicateToNodes(key, "", serverPort, config.ReplicationFactor)
	log.Printf("Data deleted and replicated for key %s", key)
	return nil
}

func (u *NodeUseCase) GetHeartbeatStatus() map[string]interface{} {
	status := u.repo.GetHeartbeatStatus()
	status["leader"] = true // Добавляем информацию о лидерстве, если это так
	return status
}

func (u *NodeUseCase) JoinCluster(address, port string) error {
	return u.repo.JoinCluster(address, port)
}

func (u *NodeUseCase) MonitorNodes() {
	for {
		time.Sleep(config.HeartbeatInterval)
		u.repo.MonitorNodes()
	}
}
