package app

//
//import (
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	v1 "github.com/pauliusluksys/golang-Reddit/handlers/v1"
//	"log"
//	"net/http"
//)
//
//func routes() *gin.Engine {
//	r := gin.New()
//	r.Use(gin.Logger())
//	r.Use(gin.Recovery())
//	apiv1 := r.Group("/api/v1")
//	apiv1.GET("/posts", v1.PostH)
//	//apiv1.GET("/:categoryName/comments/:userId/:postSlug", v1.PostCommentH)
//
//	return r
//}
//func sendJSONResponse(w http.ResponseWriter, data interface{}) {
//	body, err := json.Marshal(data)
//	if err != nil {
//		log.Printf("Failed to encode a JSON response: %v", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//	_, err = w.Write(body)
//	if err != nil {
//		log.Printf("Failed to write the response body: %v", err)
//		return
//	}
//}
