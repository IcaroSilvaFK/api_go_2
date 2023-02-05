package repositories

import "api/cmd/internals/entities"

func ShowAllTweets() map[*entities.Tweet]entities.Tweet{
	
	allTweets := entities.TweetEntity().FindAll()

	return allTweets
}