package models

import "gorm.io/gorm"

type Notes struct {
	gorm.Model
	Text string
}
