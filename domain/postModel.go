package domain

import (
	"gorm.io/gorm"
)

type PostGorm struct {
	gorm.Model
	IsDraft        bool
	DiscussionType string
	Title          string
	AuthorId       string
	CategoryId     string
	Content        string
	IsNFS          bool
	Media          string
}

//Tabler and tableName are for rewriting default association of PostGorm struct with database table "post_gorms" to "posts"
