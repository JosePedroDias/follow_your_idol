SELECT COUNT(*)
FROM twitter_tweet

SELECT user_id, COUNT(user_id)
FROM twitter_tweet
GROUP BY user_id

SELECT
    document->>'created_at' as created_at,
    document->>'text' as text
FROM twitter_tweet

----

-- "191431067" "FleurEast"

SELECT COUNT(*)
FROM twitter_user

SELECT
  screen_name,
  document->>'id' as user_id,
  document->>'name' as name
FROM twitter_user

----


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
	document->>'created_at' as created_at,
	document->'status'->>'text' as status_text
FROM twitter_user
WHERE screen_name = 'MarioMartins72';



SELECT COUNT(*)
FROM twitter_user



SELECT screen_name, COUNT(screen_name)
FROM twitter_user



SELECT screen_name AS screen_name, count(screen_name) AS count
FROM twitter_user
GROUP BY screen_name;



SELECT lang AS lang, count(lang) AS count
FROM twitter_user
GROUP BY lang;



SELECT document->>'lang' AS lang, count(document->>'lang') AS count
FROM twitter_user
GROUP BY lang
ORDER BY count DESC;



SELECT document->>'location' AS location, count(document->>'location') AS count
FROM twitter_user
GROUP BY location
ORDER BY count DESC;



SELECT document->>'time_zone' AS time_zone, count(document->>'time_zone') AS count
FROM twitter_user
GROUP BY time_zone
ORDER BY count DESC;



SELECT
	document->>'followers_count' AS followers_count,
	count(document->>'followers_count') AS count
FROM twitter_user
GROUP BY followers_count
ORDER BY count DESC



SELECT
	document->>'statuses_count' AS statuses_count,
	count(document->>'statuses_count') AS count
FROM twitter_user
GROUP BY statuses_count
ORDER BY count DESC
