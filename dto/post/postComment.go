package dtopost

import (
	"github.com/pauliusluksys/golang-Reddit/domain"
	"time"
)

type PostCommentResponse struct {
	ID        uint       `json:"post_comments_id"`
	PostId    uint       `json:"post_id"`
	AuthorId  uint       `json:"user_id"`
	ParentId  int64      `json:"parent_id,omitempty"`
	Text      string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
type PostCommentsResponse struct {
	TotalComments int
	PostComments  []PostCommentResponse
}
type PostComment struct {
	PComment domain.PostComment
}
type PostComments struct {
	PComments []domain.PostComment
}

func (pc PostComments) AllPostCommentsToDto(totalComments int) PostCommentsResponse {
	pCR := PostCommentsResponse{}
	pCR.TotalComments = totalComments
	for _, v := range pc.PComments {
		PostCommentStruct := PostComment{
			PComment: v,
		}
		postComment := PostCommentStruct.PostCommentToDto()
		pCR.PostComments = append(pCR.PostComments, postComment)
	}
	return pCR
}
func (pc PostComment) PostCommentToDto() PostCommentResponse {
	//ID        uint          `db:"post_comments_id"`
	//PostId    uint          `db:"post_id"`
	//AuthorId  uint          `db:"user_id"`
	//ParentId  sql.NullInt64 `db:"parent_id"`
	//Text      string        `db:"content"`
	//CreatedAt sql.NullTime  `db:"created_at"`
	//UpdatedAt sql.NullTime  `db:"updated_at"`
	//DeletedAt sql.NullTime  `db:"deleted_at"`
	var commentUpdatedAt *time.Time
	var commentDeletedAt *time.Time
	var ParentId int64
	if pc.PComment.UpdatedAt.Valid {
		commentUpdatedAt = &pc.PComment.UpdatedAt.Time
	}
	if pc.PComment.DeletedAt.Valid {
		commentDeletedAt = &pc.PComment.DeletedAt.Time
	}
	if pc.PComment.ParentId.Valid {
		ParentId = pc.PComment.ParentId.Int64
	}
	return PostCommentResponse{
		ID:        pc.PComment.ID,
		PostId:    pc.PComment.PostId,
		AuthorId:  pc.PComment.AuthorId,
		ParentId:  ParentId,
		Text:      pc.PComment.Text,
		CreatedAt: pc.PComment.CreatedAt.Time,
		UpdatedAt: commentUpdatedAt,
		DeletedAt: commentDeletedAt,
	}
}
