package persistence

// pgadmin
// psql -W follow_your_idol jdias
// select * from twitter_user;
// select * from twitter_tweet;

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

// ----------------------

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

func LoadUser(screenName string) (*anaconda.User, error) {
	var user anaconda.User
	var doc []byte

	row := db.QueryRow(`SELECT document
		FROM twitter_user
		WHERE screen_name = $1`, screenName)
	row.Scan(&doc)

	err := json.Unmarshal(doc, &user)
	return &user, err
}

// ----------------------

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

func LoadTweet(tweetId string) (*anaconda.Tweet, error) {
	var tweet anaconda.Tweet
	var doc []byte

	row := db.QueryRow(`SELECT document
		FROM twitter_tweet
		WHERE tweet_id = $1`, tweetId)
	row.Scan(&doc)

	err := json.Unmarshal(doc, &tweet)
	return &tweet, err
}
