package services

import "github.com/pauliusluksys/golang-Reddit/domain"

func GetAllPosts() []domain.Post {
	allPosts := domain.GetAllPosts()
	return allPosts
}
