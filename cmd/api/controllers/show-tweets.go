package controllers

import (
	use_cases "api/cmd/internals/use-cases"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func ShowAllTweets(ctx *gin.Context){


	tweets,err := use_cases.ShowTweets()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"timestamp":time.Now(),
		})
	}

	ctx.JSON(http.StatusOK,gin.H{
		"tweets": tweets,
	})
}