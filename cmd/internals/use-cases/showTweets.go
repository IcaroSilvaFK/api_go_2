package use_cases

import (
	"api/cmd/internals/entities"
)

func ShowTweets() ([]entities.Tweet, error) {

	allTweets,err := entities.FindAll()

	if err != nil {
		return nil, err
	}

	return allTweets,nil
}