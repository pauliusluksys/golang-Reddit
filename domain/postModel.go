package domain

import (
	"database/sql"
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

type PostsResponse struct {
	Posts []PostGorm `json:"post"`
}
type Post struct {
	ID             uint         `db:"id"`
	CreatedAt      sql.NullTime `db:"created_at"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
	IsDraft        bool         `db:"is_draft"`
	DiscussionType string       `db:"discussion_type"`
	Title          string       `db:"title"`
	AuthorId       string       `db:"author_id"`
	CategoryId     string       `db:"category_id"`
	Content        string       `db:"content"`
	IsNFS          bool         `db:"is_nfs"`
	Media          string       `db:"media"`
	Comments       []PostComment
}
type PostComment struct {
	ID        uint          `db:"post_comments_id"`
	PostId    uint          `db:"post_id"`
	AuthorId  uint          `db:"user_id"`
	ParentId  sql.NullInt64 `db:"parent_id"`
	Text      string        `db:"content"`
	CreatedAt sql.NullTime  `db:"created_at"`
	UpdatedAt sql.NullTime  `db:"updated_at"`
	DeletedAt sql.NullTime  `db:"deleted_at"`
}
type PostComments struct {
	Comments []PostComment
}
type PostRequest struct {
	PostId   uint   `json:"post_id"`
	UserId   uint   `json:"user_id"`
	ParentId uint   `json:"parent_id"`
	Text     string `json:"text"`
}
type PostResponse struct {
}

//Tabler and tableName are for rewriting default association of PostGorm struct with database table "post_gorms" to "posts"
