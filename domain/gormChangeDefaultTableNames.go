package domain

import (
	"github.com/pauliusluksys/golang-Reddit/domain/user"
)

type Tabler interface {
	TableName() string
}
type UserGorm struct {
	user.UserGorm
}

func (PostGorm) TableName() string {
	return "posts"
}
func (UserGorm) TableName() string {
	return "users"
}
func (PostCategoryGorm) TableName() string {
	return "categories"
}
