package use_cases

import (
	"api/cmd/internals/entities"
	"api/cmd/internals/repositories"
)

func ShowTweets() map[*entities.Tweet]entities.Tweet {

	allTweets := repositories.ShowAllTweets()

	return allTweets
}