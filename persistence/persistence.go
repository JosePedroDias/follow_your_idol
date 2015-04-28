package persistence

// pgadmin

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Setup() error {
	var err error
	db, err = sql.Open("postgres", "dbname=follow_your_idol user=jdias password=arena666 sslmode=disable")
	return err
}

func SaveUser(user *anaconda.User) error {
	doc, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO twitter_user
		(screen_name, document)
		VALUES ($1, $2)`, user.ScreenName, doc)
	return err
}

func SaveTweet(tweet *anaconda.Tweet) error {
	doc, err := json.Marshal(tweet)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO twitter_tweet
		(tweet_id, user_id, document)
		VALUES ($1, $2, $3)`, tweet.IdStr, tweet.User.IdStr, doc)
	return err
}
