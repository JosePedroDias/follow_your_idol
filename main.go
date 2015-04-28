package main

import (
	"fmt"
	"gitlab.com/josepedrodias/follow_your_idol/config"
	"gitlab.com/josepedrodias/follow_your_idol/persistence"
	"gitlab.com/josepedrodias/follow_your_idol/print"
	"gitlab.com/josepedrodias/follow_your_idol/twitter"
)

func main() {
	cfg, err := config.FromJSONFile("config.json")
	if err != nil {
		panic(err)
	}

	persistence.Setup()

	twitter.Setup(cfg)

	// users bhaenow FleurEast
	// 1st (most recent)   592717057994686464
	// last                590164133846380544
	tweets, err := twitter.GetSearch("from:bhaenow", "", "") // from:screenName, moreRecentThan, olderOrEqualTo
	if err != nil {
		panic(err)
	}

	count := len(tweets)
	fmt.Printf("\n# tweets: %d\n\n", count)

	if count > 0 {
		//print.DisplayUser()
		err := persistence.SaveUser(&tweets[0].User)
		if err != nil {
			fmt.Println("user persistence skipped")
		}
	}

	for _, tweet := range tweets {
		//fmt.Printf("#%d\n", idx)
		print.DisplayTweet(&tweet)
		err = persistence.SaveTweet(&tweet)
		if err != nil {
			fmt.Println("tweet persistence skipped")
		}
		// print.DisplayTweetAsJSON(&tweet)
	}
}
