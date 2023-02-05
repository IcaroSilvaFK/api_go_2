package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created"`
}

func NewTweet(description string) (*Tweet, error) {

	tweet := &Tweet{
		ID: uuid.NewString(),
		Description:description,
		CreatedAt: time.Now(),
	}


	err := tweet.IsValid()

	if(err != nil){
		return nil,err
	}	

	return tweet,nil

}	


func (tweet *Tweet) FindAll() []Tweet{

	tweets := []Tweet{}

	return tweets
}


func (tweet *Tweet) IsValid() error {



	fmt.Println(tweet.Description)


	if(len(tweet.Description) < 5){
		return fmt.Errorf("tweet must be grater than 5. got %d", len(tweet.Description))
	}

	return nil

}	