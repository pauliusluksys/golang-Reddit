package services

import "github.com/pauliusluksys/golang-Reddit/domain"

func GetAllPosts() []domain.PostGorm {
	allPosts := domain.GetAllPosts()
	return allPosts
}
