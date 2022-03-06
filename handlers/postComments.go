package handlers

import (
	"encoding/json"
	"fmt"
	dto "github.com/pauliusluksys/golang-Reddit/dto/post"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

func (h PostHandler) PostCommentsStore(w http.ResponseWriter, r *http.Request) {
	NewPostCommentRequest := dto.NewPostCommentRequest{}

	err := json.NewDecoder(r.Body).Decode(&NewPostCommentRequest)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "could not decode request")
	}
	newPostCommentResponse := h.PostService.NewPostComment(NewPostCommentRequest)
	utils.RespondWithJSON(w, http.StatusOK, newPostCommentResponse)
}

func (h PostHandler) PostComments(w http.ResponseWriter, r *http.Request) {
	var URLParams = make(map[string]string)
	query := r.URL.Query()
	next := query.Get("next")
	postId := query.Get("post_id")
	fmt.Println(postId)
	if next != "" {
		URLParams["next"] = next
	}
	if postId != "" {
		URLParams["postId"] = postId
	} else {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Post id not provided")
		return
	}

	postCommentsResponse := h.PostService.GetPostComments(URLParams)
	utils.RespondWithJSON(w, http.StatusOK, postCommentsResponse)
}
