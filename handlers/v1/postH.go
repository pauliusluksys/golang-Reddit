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
func PostCommentH(c *gin.Context) {
	//if err != nil {
	//	c.JSON(http., gin.H{"user": user, "value": value})
	//} else {
	//	c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//}
}
