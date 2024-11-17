package models

import "time"

type Anomaly struct {
	ID        uint   `gorm:"primaryKey"`
	SessionID string `gorm:"index"`
	Frequency float64
	Timestamp time.Time
}
