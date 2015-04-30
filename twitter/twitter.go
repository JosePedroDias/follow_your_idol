package twitter

import (
	"github.com/JosePedroDias/anaconda"
	"github.com/JosePedroDias/follow_your_idol/config"
	"net/url"
)

var api *anaconda.TwitterApi

func Setup(config *config.TwitterConfig) {
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api = anaconda.NewTwitterApi(config.AccessTokenKey, config.AccessTokenSecret)
	//api.SetDelay(10 * time.Second)
}

// https://dev.twitter.com/rest/reference/get/search/tweets

func GetSearch(query string, moreRecentThan string, olderOrEqualTo string) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("result_type", "recent") // one of: mixed, recent, popular
	v.Set("include_entities", "false")
	v.Set("count", "100") // TWEETS PER PAGE (100 is max)
	if len(olderOrEqualTo) > 0 {
		v.Set("max_id", olderOrEqualTo)
	}
	if len(moreRecentThan) > 0 {
		v.Set("since_id", moreRecentThan)
	}
	//v.Set("until", "2015-04-20")

	resp, err := api.GetSearch(query, v)
	return resp.Statuses, err
}

// https://dev.twitter.com/rest/reference/get/statuses/user_timeline
// https://dev.twitter.com/rest/public/timelines

func GetUserTimeline(screenName string, moreRecentThan string, olderOrEqualTo string) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("screen_name", screenName)
	//v.Set("trim_user", "true")
	//v.Set("exclude_replies", "true")
	//v.Set("include_rts", "true")
	v.Set("count", "200") // TWEETS PER PAGE (200 is max)
	if len(olderOrEqualTo) > 0 {
		v.Set("max_id", olderOrEqualTo)
	}
	if len(moreRecentThan) > 0 {
		v.Set("since_id", moreRecentThan)
	}
	//v.Set("until", "2015-04-20")

	tweets, err := api.GetUserTimeline(v)
	return tweets, err
}

func GetQuota() (anaconda.RateLimit, error) {
	v := url.Values{}
	rls, err := api.GetRateLimitStatus(v)
	return rls, err
}
