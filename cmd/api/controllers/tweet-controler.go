package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TweetRouter(instance *gin.Engine) {

	router := instance.Group("/_api/v1")

	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})

}