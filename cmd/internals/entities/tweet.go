package entities

import (
	"fmt"

	"github.com/google/uuid"
)

type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func TweetEntity () *Tweet{
	return &Tweet{}
}

func (tweet *Tweet) NewTweet(description string) (*Tweet, error) {

	newTweet := &Tweet{
		ID: uuid.NewString(),
		Description:description,
	}

	err := tweet.IsValid()

	if(err != nil){
		return nil,err
	}	

	return newTweet,nil

}	


func (tweet *Tweet) FindAll() map[*Tweet]Tweet{

	tweets := map[*Tweet]Tweet{}

	return tweets
}


func (tweet *Tweet) IsValid() error {

	if(len(tweet.Description) < 5){
		return fmt.Errorf("tweet must be grater than 5. got %d", len(tweet.Description))
	}

	return nil

}	