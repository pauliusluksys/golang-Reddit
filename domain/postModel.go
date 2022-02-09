package domain

import (
	"gorm.io/gorm"
)

type PostGorm struct {
	gorm.Model
}

//Tabler and tableName are for rewriting default association of PostGorm struct with database table "post_gorms" to "posts"
type Tabler interface {
	TableName() string
}

func (PostGorm) TableName() string {
	return "posts"
}
