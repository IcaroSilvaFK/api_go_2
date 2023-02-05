package use_cases

import (
	"api/cmd/internals/entities"
)

func NewCreateTweetUseCase(description string) (*entities.Tweet, error){

	tweet, err := entities.NewTweet(description)

	if err != nil {
		return nil, err
	}

	return tweet, nil

}