package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pauliusluksys/golang-Reddit/services"
	"net/http"
)

func PostH(c *gin.Context) {
	allPosts := services.GetAllPosts()
	c.JSON(http.StatusOK, gin.H{"all posts": allPosts})

}

//func PostCommentH(c *gin.Context) {
//	postSlug := c.Param("postSlug")
//	PostComments := services.GetPostComments(postSlug)
//	c.JSON(http.StatusOK, gin.H{"post_comments": allPosts})
//}
