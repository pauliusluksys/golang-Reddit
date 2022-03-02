package v1

import (
	"encoding/json"
	"fmt"
	dto "github.com/pauliusluksys/golang-Reddit/dto/post"
	"github.com/pauliusluksys/golang-Reddit/services"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

func PostCommentsUpdate(w http.ResponseWriter, r *http.Request) {

}
func PostCommentsStore(w http.ResponseWriter, r *http.Request) {
	NewPostCommentRequest := dto.NewPostCommentRequest{}

	err := json.NewDecoder(r.Body).Decode(&NewPostCommentRequest)
	if err != nil {
		fmt.Println("error:" + err.Error())
	}
	newPostCommentResponse := services.NewPostComment(NewPostCommentRequest)
	utils.RespondWithJSON(w, newPostCommentResponse)
}
func PostComments(w http.ResponseWriter, r *http.Request) {
	var URLParams = make(map[string]string)
	query := r.URL.Query()
	next := query.Get("next")
	postId := query.Get("post_id")

	if next != "" {
		URLParams["next"] = next
	}
	if postId != "" {
		URLParams["postId"] = postId
	}

	fmt.Println(postId)
	postCommentsResponse := services.GetPostComments(URLParams)
	utils.RespondWithJSON(w, postCommentsResponse)
}
