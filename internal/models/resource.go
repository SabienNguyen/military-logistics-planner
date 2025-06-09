package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model        // Provides ID, CreatedAt, UpdatedAt, DeletedAt automatically
	Type       string // Type of resource: e.g., "troop", "vehicle"
	Name       string // Human-readable name (e.g. "Alpha Squad")
	Status     string // e.g., "active", "maintenance", etc.
	ZoneID     uint   // Foreign key: links this resource to a zone
}
