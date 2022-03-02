package services

import (
	"github.com/pauliusluksys/golang-Reddit/domain"
	dto "github.com/pauliusluksys/golang-Reddit/dto/post"
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
	postCommentsDb := domain.GetPostComments(URLParams)
	totalComments := domain.GetTotalComments(URLParams)
	postCommentsStruct := dto.PostComments{postCommentsDb}

	postCommentsResponse := postCommentsStruct.AllPostCommentsToDto(totalComments)
	return postCommentsResponse
}
