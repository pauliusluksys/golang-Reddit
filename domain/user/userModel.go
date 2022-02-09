package user

import "database/sql"

type UserGorm struct {
	UserId         string
	FirstName      string
	LastName       string
	Email          string
	Password       string
	ProfilePicture string
	Country        string
	City           string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}
