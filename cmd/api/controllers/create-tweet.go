package controllers

import (
	"api/cmd/internals/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTweet(ctx  *gin.Context) {

	newTweet := repositories.NewTweet()

	ctx.JSON(http.StatusOK, gin.H{
		"tweet":newTweet,
	})
}