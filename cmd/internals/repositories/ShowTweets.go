package repositories

import "api/cmd/internals/entities"

func ShowAllTweets() []entities.Tweet{
	
	allTweets := entities.TweetEntity().FindAll()

	return allTweets
}