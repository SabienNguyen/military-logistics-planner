package models

import "time"

type Mission struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Location    string
	UnitID      uint
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
