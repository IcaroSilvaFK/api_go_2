package entities

import (
	"api/cmd/internals/database"
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
	}

	err := tweet.IsValid()
	
	if(err != nil){
		return nil,err
	}

	err = tweet.Save()


	if err != nil {
		return nil, err
	}
	return tweet,nil

}	


func FindAll()  ([]Tweet, error){

	tweets,err := listAll()

	if err != nil {
		return nil ,err
	}

	return tweets,nil
}


func (tweet *Tweet) IsValid() error {

	if(len(tweet.Description) < 5){
		return fmt.Errorf("tweet must be grater than 5. got %d", len(tweet.Description))
	}

	return nil

}	

func listAll() ([]Tweet, error){

	rows,err := database.Connection().Query(`
		SELECT * FROM tweets
	`)

	if err != nil{
		return nil, err
	}

	var tweets []Tweet
	
	for rows.Next(){
		var tweet Tweet

		rows.Scan(&tweet.ID,&tweet.Description,&tweet.CreatedAt)

		tweets = append(tweets, tweet)

	}


	return tweets,nil

}

func (tweet *Tweet) Save()  error {
	stmt,err := database.Connection().Prepare(`
		INSERT INTO tweets 
			(ID, Description)
		VALUES
		 (?,?,?)
	`)

	if err != nil {
		return err
	}	

	fmt.Println(tweet.ID)

	resp, err := stmt.Exec(tweet.ID, tweet.Description)

	if err != nil{
		return err
	}

	fmt.Println(resp.RowsAffected())

	return nil

}