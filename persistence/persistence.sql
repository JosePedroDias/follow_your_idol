CREATE TABLE twitter_tweet
(
  tweet_id character varying(32) NOT NULL,
  user_id character varying(32),
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
  document json,
  CONSTRAINT twitter_user_pkey PRIMARY KEY (screen_name)
);
