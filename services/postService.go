package services

import (
	"github.com/pauliusluksys/golang-Reddit/domain"
	dtopost "github.com/pauliusluksys/golang-Reddit/dto/post"
)

func GetAllPosts() domain.PostsResponse {
	allPosts := domain.GetAllPosts()
	postsResponse := domain.PostsResponse{Posts: allPosts}
	return postsResponse
}
func GetPost() domain.PostResponse {
	allPosts := domain.GetAllPosts()
	postsResponse := domain.PostsResponse{Posts: allPosts}
	return postsResponse
}
func GetPostComments(postSlug string) []dtopost.PostComment {
	return []dtopost.PostComment{}
}
