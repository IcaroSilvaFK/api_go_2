package repositories

import (
	"api/cmd/internals/entities"

	"github.com/google/uuid"
)

func NewTweet() *entities.Tweet {

	tweet := entities.Tweet{
		ID: uuid.NewString(),
	}

	return &tweet
}
