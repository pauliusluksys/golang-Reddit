package handlers

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/pauliusluksys/golang-Reddit/services"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

type PostHandler struct {
	PostService services.PostService
	Logger      hclog.Logger
}

func NewPostHandler(service services.PostService, logger hclog.Logger) PostHandler {
	return PostHandler{PostService: service, Logger: logger}
}
func (h PostHandler) AllPostsH(w http.ResponseWriter, r *http.Request) {
	h.Logger.Debug("get all posts handler")
	//ctxVal := r.Context()
	//user := ctxVal.Value("user_email")

	allPosts, err := h.PostService.GetAllPosts()
	if err != nil {
		err.HttpResponse(w)
	}

	utils.RespondWithOK(w, allPosts)
}
func (h PostHandler) PostH(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	postId := query.Get("post_id")
	fmt.Println(postId)
	postById, err := h.PostService.GetPost(postId)
	if err != nil {
		err.HttpResponse(w)
	}
	utils.RespondWithOK(w, postById)
}

//func PostCommentH(c *gin.Context) {
//	postSlug := c.Param("postSlug")
//	PostComments := services.GetPostComments(postSlug)
//	c.JSON(http.StatusOK, gin.H{"post_comments": allPosts})
//}
