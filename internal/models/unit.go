package models

import "time"

type Unit struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Branch        string
	Location      string
	CommanderName string
	ContactEmail  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
