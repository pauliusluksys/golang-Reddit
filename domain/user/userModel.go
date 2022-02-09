package user

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type UserGorm struct {
	gorm.Model
	FirstName      string
	LastName       string
	Email          string
	Password       string
	ProfilePicture string
	Country        string
	City           string
}

type User struct {
	ID             uint         `db:"id"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
	FirstName      string       `db:"first_name"`
	LastName       string       `db:"last_name"`
	Email          sql.NullString
	Password       string
	ProfilePicture sql.NullString `db:"profile_picture"`
	Country        sql.NullString
	City           sql.NullString
}
