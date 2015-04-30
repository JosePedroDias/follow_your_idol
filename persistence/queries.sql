-- http://www.postgresql.org/docs/9.3/static/functions.html
-- http://www.postgresql.org/docs/9.3/static/functions-aggregate.html

SELECT COUNT(*)
FROM twitter_tweet;

SELECT
	user_id,
	COUNT(user_id) as num_tweets,
	MIN(TO_TIMESTAMP(document->>'created_at', 'Dy Mon DD HH24:MI:SS 9999 YYYY')) as oldest,
	MAX(TO_TIMESTAMP(document->>'created_at', 'Dy Mon DD HH24:MI:SS 9999 YYYY')) as newest,
	MIN(tweet_id) as oldest_id,
	MAX(tweet_id) as newest_id
FROM twitter_tweet
GROUP BY user_id;

SELECT
	t.user_id,
	COUNT(t.user_id) as num_tweets,
	MIN(TO_TIMESTAMP(t.document->>'created_at', 'Dy Mon DD HH24:MI:SS 9999 YYYY')) as oldest,
	MAX(TO_TIMESTAMP(t.document->>'created_at', 'Dy Mon DD HH24:MI:SS 9999 YYYY')) as newest,
	MIN(t.tweet_id) as oldest_id,
	MAX(t.tweet_id) as newest_id
FROM twitter_tweet as t, twitter_user as u
WHERE u.user_id = t.user_id
GROUP BY t.user_id;


SELECT
    tweet_id,
    TO_TIMESTAMP(document->>'created_at', 'Dy Mon DD HH24:MI:SS 9999 YYYY') as created_at,
    document->'user'->>'screen_name' as from,
    document->>'text' as text,
    document->>'favorite_count' as favorites,
    document->>'retweet_count' as retweets
FROM twitter_tweet
where tweet_id = '592717057994686464';

----

-- 191431067 'FleurEast'
-- 702764472 'Bhaenow'

SELECT COUNT(*)
FROM twitter_user;

SELECT
  screen_name,
  user_id
FROM twitter_user;

SELECT
	document->>'screen_name' as screen_name,
	document->>'id_str' as user_id,
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

select screen_name, document->>'id' as user_id from twitter_user;
