package models

import "time"

type SupplyRequest struct {
	ID          uint `gorm:"primaryKey"`
	UnitID      uint
	DepotID     uint
	MissionID   uint
	ItemName    string
	Quantity    uint
	Status      string
	RequestedAt time.Time
	FulfilledAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
