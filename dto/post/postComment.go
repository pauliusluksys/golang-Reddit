package dtopost

import (
	"time"
)

type NewPostCommentRequest struct {
	PostId   uint   `json:"post_id"`
	AuthorId uint   `json:"user_id"`
	ParentId int64  `json:"parent_id"`
	Content  string `json:"content"`
}
type NewPostCommentResponse struct {
	ID        uint      `json:"post_comments_id"`
	PostId    uint      `json:"post_id"`
	AuthorId  uint      `json:"user_id"`
	ParentId  int64     `json:"parent_id,omitempty"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
type PostCommentResponse struct {
	ID        uint       `json:"post_comments_id"`
	PostId    uint       `json:"post_id"`
	AuthorId  uint       `json:"user_id"`
	ParentId  int64      `json:"parent_id,omitempty"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
type PostCommentsResponse struct {
	TotalComments int
	PostComments  []PostCommentResponse
}
