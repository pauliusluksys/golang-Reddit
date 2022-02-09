package domain

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Code  string
	Price uint
}
