package domain

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/pauliusluksys/golang-Reddit/errs"
	"gorm.io/gorm"
)

type PostGormRepo struct {
	client *gorm.DB
	Logger hclog.Logger
}
type PostSqlxRepo struct {
	client *sqlx.DB
	Logger hclog.Logger
}

func NewPostSqlxRepo(db *sqlx.DB, logger hclog.Logger) PostSqlxRepo {
	return PostSqlxRepo{client: db, Logger: logger}
}
func NewPostGormRepoDb(gormDb *gorm.DB, logger hclog.Logger) PostGormRepo {
	return PostGormRepo{client: gormDb, Logger: logger}
}
func (psqlx PostSqlxRepo) GetAllPosts() (*[]Post, *errs.AppError) {

	var posts []Post
	query := "SELECT * FROM posts LIMIT 20;"
	err := psqlx.client.Select(&posts, query)
	if err != nil {
		return nil, errs.NewNotFoundError("Something went wrong when querying for all posts")
	}

	return &posts, nil
}
func (psqlx PostSqlxRepo) GetPostById(postId string) (*Post, *errs.AppError) {
	post := []Post{}
	fmt.Println(postId)
	query := "SELECT * FROM posts where id=? LIMIT 1;"
	err := psqlx.client.Select(&post, query, postId)
	if err != nil {
		return nil, errs.NewUnexpectError("Something went wrong when querying post by id")
	}
	fmt.Println(post)
	return &post[0], nil
}
