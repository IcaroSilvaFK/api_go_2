package main

import (
	"api/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {

	router := engine.Group("/_api/v1")

	router.GET("/ping",controllers.CreateTweet)

}