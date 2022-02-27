package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
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
type PostsResponse struct {
	Posts []PostGorm `json:"post"`
}
type Post struct {
	ID             uint         `db:"id"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      time.Time    `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
	IsDraft        bool         `db:"is_draft"`
	DiscussionType string       `db:"discussion_type"`
	Title          string       `db:"title"`
	AuthorId       string       `db:"author_id"`
	CategoryId     string       `db:"category_id"`
	Content        string       `db:"content"`
	IsNFS          bool         `db:"is_nfs"`
	Media          string       `db:"media"`
}
type PostComment struct {
	ID        uint         `db:"id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	PostId    uint         `db:"post_id"`
	AuthorId  uint         `db:"author_id"`
	ParentId  uint         `db:"parent_id"`
	Text      string       `db:"text"`
}
type PostResponse struct {
}

//Tabler and tableName are for rewriting default association of PostGorm struct with database table "post_gorms" to "posts"
