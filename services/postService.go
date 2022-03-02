package services

import (
	"database/sql"
	"github.com/pauliusluksys/golang-Reddit/domain"
	dto "github.com/pauliusluksys/golang-Reddit/dto/post"
	"time"
)

func GetAllPosts() domain.PostsResponse {
	allPosts := domain.GetAllPosts()
	postsResponse := domain.PostsResponse{Posts: allPosts}
	return postsResponse
}
func GetPost(postId string) domain.Post {
	postById := domain.GetPostById(postId)
	postsResponse := postById
	return postsResponse
}
func GetPostComments(URLParams map[string]string) dto.PostCommentsResponse {
	postComments := domain.GetPostComments(URLParams)
	totalComments := domain.GetTotalComments(URLParams)
	allPostComments := domain.PostComments{postComments}
	postCommentsResponse := allPostComments.AllPostCommentsToDto(totalComments)
	return postCommentsResponse
}
func NewPostComment(newPostCommentReq dto.NewPostCommentRequest) dto.NewPostCommentResponse {
	newPostComment := domain.PostComment{
		PostId:    newPostCommentReq.PostId,
		AuthorId:  newPostCommentReq.AuthorId,
		Content:   newPostCommentReq.Content,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	if newPostCommentReq.ParentId == 0 {
		newPostComment.ParentId = sql.NullInt64{Valid: false}
	} else {
		newPostComment.ParentId = sql.NullInt64{Int64: newPostCommentReq.ParentId, Valid: true}
	}
	newPostComment = domain.NewPostComment(newPostComment)
	postCommentResponse := newPostComment.NewPostCommentToDto()
	return postCommentResponse
}
