package repositories

import "api/cmd/internals/entities"

func ShowAllTweets() []entities.Tweet{
	
	allTweets := entities.FindAll()

	return allTweets
}