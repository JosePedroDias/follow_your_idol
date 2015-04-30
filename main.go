package main

import (
	"flag"
	"fmt"
	"github.com/JosePedroDias/anaconda" // ChimeraCoder JosePedroDias
	"gitlab.com/josepedrodias/follow_your_idol/config"
	"gitlab.com/josepedrodias/follow_your_idol/persistence"
	"gitlab.com/josepedrodias/follow_your_idol/print"
	"gitlab.com/josepedrodias/follow_your_idol/twitter"
	"time"
)

const ( // Mon Jan 2 15:04:05 MST 2006
	HHMMSS = "15:04:05"
)

func main() {

	// 0) READ CONFIG AND PREPARE MODULES
	tCfg, err := config.TwitterFromFile("twitter_config.json")
	if err != nil {
		panic(err)
	}
	pCfg, err := config.PostgresqlFromFile("postgresql_config.json")
	if err != nil {
		panic(err)
	}

	// https://golang.org/pkg/flag/
	// https://gobyexample.com/command-line-flags
	cmdTimeline := flag.String("timeline", "", `command. fetches user timeline tweets using the twitter API. Accepts screenName ex:"Bhaenow"`)
	cmdSearch := flag.String("search", "", `command. searches tweets using the twitter API. Accepts query string. ex:"from:Bhaenow"`)
	cmdQuota := flag.Bool("quota", false, `command. returns relevant quotas using the twitter API.  Does not accept argument.`)

	cmdDbStats := flag.Bool("db_stats", false, `command. returns database status. Does not accept argument.`)
	cmdGetTweet := flag.String("get_tweet", "", `command. returns cached tweet from the database. Accepts twitterId. ex:"592717057994686464"`)
	cmdGetUser := flag.String("get_user", "", `command. returns cached twitter user from the database. Accepts screenName. ex:"Bhaenow"`)

	optMoreRecentThan := flag.String("more_recent_than", "", `option. if passed, filters only tweets more recent than given tweet id. ex:"592717057994686464"`)
	optOlderOrEqualTo := flag.String("older_or_equal_to", "", `option. if passed, filters only tweets older or equal to given tweet id. ex:"592717057994686464"`)

	flag.Parse()

	fmt.Println("\nParsed arguments:")
	fmt.Printf("timeline  [%v]\n", *cmdTimeline)
	fmt.Printf("search    [%v]\n", *cmdSearch)
	fmt.Printf("quota     [%v]\n", *cmdQuota)

	fmt.Printf("db_stats  [%v]\n", *cmdDbStats)
	fmt.Printf("get_tweet [%v]\n", *cmdGetTweet)
	fmt.Printf("get_user  [%v]\n", *cmdGetUser)
	fmt.Println("----")
	fmt.Printf("more_recent_than  [%v]\n", *optMoreRecentThan)
	fmt.Printf("older_or_equal_to [%v]\n", *optOlderOrEqualTo)
	fmt.Println("")

	persistence.Setup(pCfg)
	twitter.Setup(tCfg)

	if *cmdQuota {
		quota, err := twitter.GetQuota()
		if err != nil {
			panic(err)
		}
		utLim := quota.Resources.Statuses.UserTimeline
		seLim := quota.Resources.Search.Tweets
		fmt.Println(" service  | quota     | reset")
		fmt.Println("----------+-----------+----------")
		fmt.Printf(" search   | %3d / %3d | %s\n", utLim.Remaining, utLim.Limit, time.Unix(utLim.Reset, 0).Format(HHMMSS))
		fmt.Printf(" timeline | %3d / %3d | %s\n", seLim.Remaining, seLim.Limit, time.Unix(seLim.Reset, 0).Format(HHMMSS))
	} else if *cmdDbStats {
		statuses, err := persistence.GetTwitterUserStatus()
		if err != nil {
			panic(err)
		}
		fmt.Println(" user_id    | screen_name     | tweets | oldest     | newest     | oldest_id          | newest_id")
		fmt.Println("------------+-----------------+--------+------------+------------+--------------------+--------------------")
		for _, s := range statuses {
			fmt.Printf(" %-10s | %-15s | %6d | %s | %s | %s | %s\n", s.UserId, s.ScreenName, s.NumTweets, s.Oldest.Format("2006-01-02"), s.Newest.Format("2006-01-02"), s.OldestId, s.NewestId)
		}
	} else if len(*cmdTimeline) > 0 || len(*cmdSearch) > 0 {
		var tweets []anaconda.Tweet
		var err error

		if len(*cmdTimeline) > 0 {
			tweets, err = twitter.GetUserTimeline(*cmdTimeline, *optMoreRecentThan, *optOlderOrEqualTo)
		} else {
			tweets, err = twitter.GetSearch(*cmdSearch, *optMoreRecentThan, *optOlderOrEqualTo)
		}

		if err != nil {
			panic(err)
		}

		count := len(tweets)
		stored := 0

		if count > 0 {
			//print.DisplayUser()
			err := persistence.SaveUser(&tweets[0].User)
			if err != nil {
				fmt.Println("user persistence skipped")
			}
		}

		for _, tweet := range tweets {
			print.DisplayTweet(&tweet)
			//print.DisplayTweetAsJSON(&tweet)

			err = persistence.SaveTweet(&tweet)
			if err != nil {
				fmt.Println("tweet persistence skipped")
			} else {
				stored += 1
			}
		}

		fmt.Printf("\n# tweets: %d / %d\n\n", stored, count)
	} else if len(*cmdGetTweet) > 0 {
		tweet, _ := persistence.LoadTweet(*cmdGetTweet)
		print.DisplayTweet(tweet)
	} else if len(*cmdGetUser) > 0 {
		user, _ := persistence.LoadUser(*cmdGetUser)
		print.DisplayUser(user)
	} else {
		fmt.Println("No command passed - nothing to do. Use -help to learn the command-line API.")
	}
}
