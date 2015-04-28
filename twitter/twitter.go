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

func GetSearch(query string) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("result_type", "recent") // mixed, recent, popular
	v.Set("include_entities", "false")
	v.Set("count", "100") // TWEETS PER PAGE
	//v.Set("since_id", "2015-04-14") // MORE RECENT THAN
	// v.Set("max_id", "590532726979198977") // OLDER OR EQUAL TO
	tweets, err := api.GetSearch(query, v)
	return tweets, err
}
