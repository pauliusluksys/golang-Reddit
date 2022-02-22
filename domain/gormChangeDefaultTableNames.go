package domain

type Tabler interface {
	TableName() string
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
