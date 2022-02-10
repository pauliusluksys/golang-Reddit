package services

import (
	"github.com/pauliusluksys/golang-Reddit/domain"
	dtopost "github.com/pauliusluksys/golang-Reddit/dto/post"
)

func GetAllPosts() []domain.PostGorm {
	allPosts := domain.GetAllPosts()
	return allPosts
}
func GetPostComments(postSlug string) []dtopost.PostComment {
	return []dtopost.PostComment{}
}
