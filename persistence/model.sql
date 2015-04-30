CREATE TABLE twitter_tweet
(
  tweet_id character varying(32) NOT NULL,
  user_id character varying(32) NOT NULL,
  document json,
  CONSTRAINT twitter_tweet_pkey PRIMARY KEY (tweet_id)
);

CREATE INDEX twitter_tweet_user_id_idx
  ON twitter_tweet
  USING btree
  (user_id COLLATE pg_catalog."default");

CREATE TABLE twitter_user
(
  screen_name character varying(15) NOT NULL,
  user_id character varying(32) NOT NULL,
  document json,
  CONSTRAINT twitter_user_pkey PRIMARY KEY (screen_name)
);

-- FOR FULL TEXT SEARCH

CREATE MATERIALIZED VIEW twitter_tweet_indexed AS 
SELECT
  user_id,
  tweet_id,
  document->>'text' as body,
  to_tsvector('english', document->>'text') as body2
FROM twitter_tweet;

-- CREATE INDEX

CREATE INDEX idx_body2_search ON twitter_tweet_indexed USING gin(body2);
