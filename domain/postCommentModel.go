package domain

import "database/sql"

type Comments struct {
	ID        uint         `db:"id"`
	Comment   string       `db:"comment"`
	ParentId  uint         `db:"parent_id"`
	PostId    uint         `db:"post_id"`
	UserId    uint         `db:"user_id"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
