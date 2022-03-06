package services

import (
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/pauliusluksys/golang-Reddit/domain"
	dto "github.com/pauliusluksys/golang-Reddit/dto/post"
	"github.com/pauliusluksys/golang-Reddit/errs"
	"time"
)

type PostService struct {
	PostServiceSqlx domain.PostSqlxRepo
	PostServiceGorm domain.PostGormRepo
	Logger          hclog.Logger
}

func NewPostService(psqlxRepo domain.PostSqlxRepo, logger hclog.Logger) PostService {
	return PostService{PostServiceSqlx: psqlxRepo, Logger: logger}
}

func (ps PostService) GetAllPosts() (*domain.PostsResponse, *errs.AppError) {
	ps.Logger.Debug("get all posts func")
	allPosts, err := ps.PostServiceSqlx.GetAllPosts()
	if err != nil {
		return nil, err
	}
	postsResponse := domain.PostsResponse{Posts: *allPosts}

	return &postsResponse, err
}
func (ps PostService) GetPost(postId string) (*domain.Post, *errs.AppError) {
	postById, err := ps.PostServiceSqlx.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	postsResponse := postById
	return postsResponse, err
}
func (ps PostService) GetPostComments(URLParams map[string]string) dto.PostCommentsResponse {
	postComments := ps.PostServiceSqlx.GetPostComments(URLParams)
	totalComments := ps.PostServiceSqlx.GetTotalComments(URLParams)
	allPostComments := domain.PostComments{Comments: postComments}
	postCommentsResponse := allPostComments.AllPostCommentsToDto(totalComments)
	return postCommentsResponse
}
func (ps PostService) NewPostComment(newPostCommentReq dto.NewPostCommentRequest) dto.NewPostCommentResponse {
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
	newPostComment = ps.PostServiceSqlx.NewPostComment(newPostComment)
	postCommentResponse := newPostComment.NewPostCommentToDto()
	return postCommentResponse
}
