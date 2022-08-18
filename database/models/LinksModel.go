package models

import "gorm.io/gorm"

type Links struct {
	gorm.Model
	Url  string
	Name string
}
