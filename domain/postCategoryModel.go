package domain

import "gorm.io/gorm"

type PostCategoryGorm struct {
	gorm.Model
	styles      string
	Name        string
	Subscribers int
	Title       string
	Type        string
	Path        string
}

//Id           int
//Styles       string
//Name         string
//Subscribers  int
//Title        string
//Type         string
//Path         string
//IsFavorite   bool
//IsNSFW       bool
//IsSubscribed bool
//IsEnabled    bool
