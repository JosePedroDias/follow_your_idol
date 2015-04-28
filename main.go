package main

import (
	"fmt"
	"gitlab.com/josepedrodias/follow_your_idol/config"
	"gitlab.com/josepedrodias/follow_your_idol/persistence"
	// "gitlab.com/josepedrodias/follow_your_idol/print"
	"gitlab.com/josepedrodias/follow_your_idol/twitter"
)

func main() {
	cfg, err := config.FromJSONFile("config.json")
	if err != nil {
		panic(err)
	}

	persistence.Setup()

	twitter.Setup(cfg)

	tweets, err := twitter.GetSearch("from:FleurEast") // bhaenow FleurEast
	if err != nil {
		panic(err)
	}

	count := len(tweets)
	fmt.Printf("count: %d\n", count)

	if count > 0 {
		//print.DisplayUser()
		persistence.SaveUser(&tweets[0].User)
	}

	/*for _, tweet := range tweets {
		//fmt.Printf("#%d\n", idx)
		print.DisplayTweet(&tweet)
		err = persistence.SaveTweet(&tweet)
		if err != nil {
			panic(err)
		}
		// print.DisplayTweetAsJSON(&tweet)
	}*/
}
