package entity

import "time"

type Node struct {
	Address       string    `json:"address"`
	Port          string    `json:"port"`
	IsLeader      bool      `json:"is_leader"`
	LastHeartbeat time.Time `json:"last_heartbeat"`
}
