package controllers

import (
	use_cases "api/cmd/internals/use-cases"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


type createTweetInput struct {
	Description string `json:"description" binding:"required"`
}

func CreateTweet(ctx  *gin.Context) {

	var body createTweetInput;


	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":"Description is required for creating a new tweet",
			"timestamp":time.Now(),
		})

		return
	}

	newTweet, err := use_cases.NewCreateTweetUseCase(body.Description)

	if err != nil {
		 ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"error": err.Error(),
			"timestamp":time.Now(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"tweet":newTweet,
		"timestamp":time.Now(),
	})
}
