package models

import "gorm.io/gorm"

type Zone struct {
	gorm.Model
	Name        string
	Description string
	Location    string
}
