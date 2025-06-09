package models

import "time"

type Depot struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Location     string
	Capacity     uint
	ManagerName  string
	ContactEmail string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
