package app

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pauliusluksys/golang-Reddit/handlers/v1"
)

func routes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/posts", v1.PostH)
	apiv1.GET("/comments", v1.PostCommentH)

	return r
}
