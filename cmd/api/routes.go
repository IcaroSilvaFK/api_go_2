package main

import (
	"api/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {

	router := engine.Group("/_api/v1")

	router.POST("/tweets",controllers.CreateTweet)
	router.GET("/tweets",controllers.ShowAllTweets)

}