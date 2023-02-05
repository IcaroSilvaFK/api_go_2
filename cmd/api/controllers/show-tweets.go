package controllers

import (
	"api/cmd/internals/entities"
	// "api/cmd/internals/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func ShowAllTweets(ctx *gin.Context){


	// allTweets := repositories.ShowAllTweets()

	ctx.JSON(http.StatusOK,gin.H{
		"tweets": nil,
	})
}