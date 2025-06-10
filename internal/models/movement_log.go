package models

import "gorm.io/gorm"

type MovementLog struct {
	gorm.Model
	ResourceID uint
	FromZoneID *uint
	ToZoneID   uint
	Note       string
}
