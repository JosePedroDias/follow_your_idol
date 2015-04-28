SELECT COUNT(*)
FROM twitter_tweet;

SELECT user_id, COUNT(user_id)
FROM twitter_tweet
GROUP BY user_id;

SELECT
    document->>'created_at' as created_at,
    tweet_id,
    document->>'text' as text,
    document->>'favorite_count' as favorites,
    document->>'retweet_count' as retweets
FROM twitter_tweet;

----

-- 191431067 'FleurEast'
-- 702764472 'Bhaenow'

SELECT COUNT(*)
FROM twitter_user;

SELECT
  document->>'id' as user_id,
  screen_name,
  document->>'name' as name
FROM twitter_user;

SELECT
	document->>'screen_name' as screen_name,
	document->>'name' as name,
	document->>'description' as description,
	document->>'lang' as lang,
	document->>'time_zone' as time_zone,
	document->>'location' as location,
	document->>'lang' as lang,
	document->>'followers_count' as followers_count,
	document->>'statuses_count' as statuses_count,
	document->>'created_at' as created_at
FROM twitter_user
WHERE screen_name = 'Bhaenow';
