package main

import (
	"fmt"
	"gitlab.com/josepedrodias/follow_your_idol/config"
	"gitlab.com/josepedrodias/follow_your_idol/persistence"
	"gitlab.com/josepedrodias/follow_your_idol/print"
	"gitlab.com/josepedrodias/follow_your_idol/twitter"
)

func main() {
	tCfg, err := config.TwitterFromFile("twitter_config.json")
	if err != nil {
		panic(err)
	}
	pCfg, err := config.PostgresqlFromFile("postgresql_config.json")
	if err != nil {
		panic(err)
	}

	persistence.Setup(pCfg)

	twitter.Setup(tCfg)

	// 1) FETCH TWEETS FROM USER AND PERSIST THEM

	// users Bhaenow FleurEast AndreaFaustini1
	// 1st (most recent)   592717057994686464
	// last                590164133846380544
	//tweets, err := twitter.GetSearch("from:AndreaFaustini1", "", "") // from:screenName, moreRecentThan, olderOrEqualTo
	tweets, err := twitter.GetUserTimeline("FleurEast", "", "") // from:screenName, moreRecentThan, olderOrEqualTo
	if err != nil {
		panic(err)
	}

	count := len(tweets)

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

	fmt.Printf("\n# tweets: %d\n\n", count)

	// 2) DISPLAY CACHED USER

	//user, _ := persistence.LoadUser("Bhaenow")
	//print.DisplayUser(user)

	// 3) DISPLAY CACHED TWEET

	//tweet, _ := persistence.LoadTweet("592717057994686464")
	//print.DisplayTweet(tweet)
}
