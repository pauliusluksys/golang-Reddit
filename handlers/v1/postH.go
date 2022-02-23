package v1

import (
	"fmt"
	"github.com/pauliusluksys/golang-Reddit/services"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

func PostH(w http.ResponseWriter, r *http.Request) {
	ctxVal := r.Context()
	user := ctxVal.Value("user_email")
	fmt.Println(user)
	allPosts := services.GetAllPosts()
	utils.RespondWithJSON(w, allPosts)
}

//func PostCommentH(c *gin.Context) {
//	postSlug := c.Param("postSlug")
//	PostComments := services.GetPostComments(postSlug)
//	c.JSON(http.StatusOK, gin.H{"post_comments": allPosts})
//}
