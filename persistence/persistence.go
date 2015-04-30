package persistence

// pgadmin
// psql -W follow_your_idol jdias
// select * from twitter_user;
// select * from twitter_tweet;

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/JosePedroDias/anaconda"
	"github.com/JosePedroDias/follow_your_idol/config"
	_ "github.com/lib/pq"
	"time"
)

type TwitterUserStatus struct {
	UserId     string    `json:"user_id"`
	ScreenName string    `json:"screen_name"`
	NumTweets  int       `json:"num_tweets"`
	Oldest     time.Time `json:"oldest"`
	Newest     time.Time `json:"newest"`
	OldestId   string    `json:"oldest_id"`
	NewestId   string    `json:"oldest_id"`
}

type TwitterTweetFTSResult struct {
	UserId  string `json:"user_id"`
	TweetId string `json:"tweet_id"`
	Body    string `json:"body"`
}

var db *sql.DB

func Setup(config *config.PostgresqlConfig) error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s", config.DbName, config.User, config.Password, config.SslMode))
	return err
}

// ----------------------

func SaveUser(user *anaconda.User) error {
	doc, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO twitter_user
		(screen_name, user_id, document)
		VALUES ($1, $2, $3)`, user.ScreenName, user.IdStr, doc)
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

func screenNameToUserId(screenName string) (string, error) {
	var userId string
	row := db.QueryRow(`SELECT
		user_id
	FROM twitter_user
	WHERE screen_name = $1`, screenName)
	err := row.Scan(&userId)
	return userId, err
}

func GetUserIdToScreenNameMap() (map[string]string, error) {
	resMap := make(map[string]string)

	rows, err := db.Query(`SELECT
		user_id,
		screen_name
	FROM twitter_user`)
	if err != nil {
		return resMap, err
	}
	defer rows.Close()

	for rows.Next() {
		var userId, screenName string
		if err := rows.Scan(&userId, &screenName); err != nil {
			return resMap, err
		}
		resMap[userId] = screenName
	}

	if err := rows.Err(); err != nil {
		return resMap, err
	}

	return resMap, err
}

func GetTwitterUserStatus() ([]TwitterUserStatus, error) {
	var results []TwitterUserStatus = make([]TwitterUserStatus, 0)

	userIdToScreenNameMap, err := GetUserIdToScreenNameMap()
	if err != nil {
		return results, err
	}

	rows, err := db.Query(`SELECT
	user_id,
	COUNT(user_id) as num_tweets,
	MIN(TO_TIMESTAMP(document->>'created_at', 'Dy Mon DD HH24:MI:SS 9999 YYYY')) as oldest,
	MAX(TO_TIMESTAMP(document->>'created_at', 'Dy Mon DD HH24:MI:SS 9999 YYYY')) as newest,
	MIN(tweet_id) as oldest_id,
	MAX(tweet_id) as newest_id
FROM twitter_tweet
GROUP BY user_id`)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var tus TwitterUserStatus
		if err := rows.Scan(&tus.UserId, &tus.NumTweets, &tus.Oldest, &tus.Newest, &tus.OldestId, &tus.NewestId); err != nil {
			return results, err
		}
		tus.ScreenName = userIdToScreenNameMap[tus.UserId]
		results = append(results, tus)
	}

	if err := rows.Err(); err != nil {
		return results, err
	}

	return results, err
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

func GetEdgeTweetsForUser(screenName string, isMax bool) (string, error) {
	userId, err := screenNameToUserId(screenName)
	if err != nil {
		return "", err
	}

	var query, tweetId string
	if isMax {
		query = `SELECT MAX(tweet_id) FROM twitter_tweet WHERE user_id = $1 GROUP BY user_id`
	} else {
		query = `SELECT MIN(tweet_id) FROM twitter_tweet WHERE user_id = $1 GROUP BY user_id`
	}
	row := db.QueryRow(query, userId)
	err = row.Scan(&tweetId)

	return tweetId, err
}

// ----------------------

func FullTextSearchTweets(query string) ([]TwitterTweetFTSResult, error) {
	var results []TwitterTweetFTSResult = make([]TwitterTweetFTSResult, 0)

	rows, err := db.Query(`SELECT
  user_id,
  tweet_id,
  body
FROM twitter_tweet_indexed
WHERE body2 @@ to_tsquery('english', $1)
ORDER BY ts_rank(body2, to_tsquery('english', $1)) DESC`, query)
	defer rows.Close()

	for rows.Next() {
		var r TwitterTweetFTSResult
		if err := rows.Scan(&r.UserId, &r.TweetId, &r.Body); err != nil {
			return results, err
		}
		results = append(results, r)
	}

	if err := rows.Err(); err != nil {
		return results, err
	}

	return results, err
}
