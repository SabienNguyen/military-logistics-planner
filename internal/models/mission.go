// models/mission.go
package models

import (
	"time"
)

type Mission struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Status      string
}
