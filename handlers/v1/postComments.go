package v1

import (
	"fmt"
	"github.com/pauliusluksys/golang-Reddit/services"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

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
