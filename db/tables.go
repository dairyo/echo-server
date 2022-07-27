package db

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Body string
}
