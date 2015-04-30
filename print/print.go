package print

import (
	"encoding/json"
	"fmt"
	"github.com/JosePedroDias/anaconda"
)

func DisplayTweet(tweet *anaconda.Tweet) {
	//fmt.Printf("--( TWEET )--\n")
	fmt.Printf("\n----\n\n")
	fmt.Printf("tweet id:         %s\n", tweet.IdStr)
	fmt.Printf("%s (@%s)\n", tweet.User.Name, tweet.User.ScreenName)
	fmt.Printf("user id:         %s\n", tweet.User.IdStr)
	fmt.Printf("created at: %s\n", tweet.CreatedAt)
	//fmt.Printf("created at2: %s\n", tweet.CreatedAtTime())
	//fmt.Printf("place:      %v\n", tweet.Place.Country)
	//fmt.Printf("geo: %v\n", tweet.Geo)
	fmt.Printf("text:       %s\n", tweet.Text)
	//fmt.Printf("retweeted:  %t\n", tweet.Retweeted)
	fmt.Printf("retweeted:  %d\n", tweet.RetweetCount)
	//tweet.Entities.Hashtags
	//tweet.Entities.Media
	//tweet.Entities.Urls
	//tweet.Entities.User_mentions
}

func DisplayTweetAsJSON(tweet *anaconda.Tweet) {
	//tweetJSON, err := json.Marshal(tweet)
	tweetJSON, err := json.MarshalIndent(tweet, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n----\n\n")
	fmt.Printf("%s\n", tweetJSON)
}

func DisplayUser(user *anaconda.User) {
	//fmt.Printf("--( USER )--\n")
	fmt.Printf("name:        %s\n", user.Name)
	fmt.Printf("desc:        %s\n", user.Description)
	fmt.Printf("id:          %d\n", user.Id)
	fmt.Printf("lang:        %s\n", user.Lang)
	fmt.Printf("screen name: %s\n", user.ScreenName)
	fmt.Printf("time zone:   %s\n", user.TimeZone)
	fmt.Printf("# status:    %d\n", user.StatusesCount)
	fmt.Printf("# followers: %d\n", user.FollowersCount)
	fmt.Printf("location:    %s\n", user.Location)
	fmt.Printf("created at:  %s\n", user.CreatedAt)

	/*if user.Status != nil {
		DisplayTweet(user.Status)
	}*/
}
