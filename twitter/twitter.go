package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"gitlab.com/josepedrodias/follow_your_idol/config"
	"net/url"
)

var api *anaconda.TwitterApi

func Setup(config *config.Config) {
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api = anaconda.NewTwitterApi(config.AccessTokenKey, config.AccessTokenSecret)
	//api.SetDelay(10 * time.Second)
}

// https://dev.twitter.com/rest/reference/get/search/tweets

func GetSearch(query string, moreRecentThan string, olderOrEqualTo string) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("result_type", "recent") // mixed, recent, popular
	v.Set("include_entities", "false")
	v.Set("count", "100") // TWEETS PER PAGE
	if len(olderOrEqualTo) > 0 {
		v.Set("max_id", olderOrEqualTo)
	}
	if len(moreRecentThan) > 0 {
		v.Set("since_id", moreRecentThan)
	}
	//v.Set("until", "2015-04-20")

	tweets, err := api.GetSearch(query, v)
	return tweets, err
}
