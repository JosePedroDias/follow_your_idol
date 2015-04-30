# what is this for?

The purpose of this command-line tool is to allow you scrap tweets using the Twitter API to a local PostgreSQL database.
Then one can query users and tweets locally, update timelines and do full text searches.

You must have a recent version of Go (1.4 used here).
This project uses PortgreSQL JSON columns to store tweets and users so the PostgreSQL instance must be at least 9.3


# setup

Create the files `postgresql_config.json` and `postgresql_config.json`.  
You can use the example files `postgresql_config_example.json` and `postgresql_config_example.json` as reference for the required fields.


# how to use

    go build
    ./follow_your_idol [-<op arg>]* [<-flag>]*

use `./follow_your_idol -help` for an updated list of operations and flags

Some operations work with the twitter API, some others make exclusive work of the data cached in your database.

these use twitter: `timeline`, `search`, `quota`
these don't: `db_stats` `fts`, `get_tweet`, `get_user`

these are optional, in both `timeline` and `search` operations: `more_recent_than`, `older_or_equal_to` (for twitter pagination, supports a twitter_id string or `auto`)


## examples

	> ./follow_your_idol -quota

	 service  | quota     | reset
	----------+-----------+----------
	 search   | 180 / 180 | 18:55:55
	 timeline | 180 / 180 | 18:55:55

----

	./follow_your_idol -db_stats

	 user_id    | screen_name     | tweets | oldest     | newest     | oldest_id          | newest_id
	------------+-----------------+--------+------------+------------+--------------------+--------------------
	 21329785   | JAYJAMES        |   2556 | 2014-11-02 | 2015-04-30 | 529009325881425920 | 593648206879293440
	 489758148  | AndreaFaustini1 |    854 | 2014-02-17 | 2015-04-30 | 435513101421916160 | 593701937096327169
	 1901060970 | jwaltonmusic    |    867 | 2013-09-24 | 2015-04-29 | 382552671506661376 | 593470899493457920
	 1343977518 | OverloadGen     |   3168 | 2014-08-27 | 2015-03-19 | 504762550782001154 | 578625173944987648
	 2788994607 | StereoKicks     |   3234 | 2014-12-05 | 2015-04-30 | 540814293584076800 | 593706883040555008
	 2739844841 | THESTEVIRITCHIE |   3205 | 2014-10-03 | 2015-04-29 | 518081464236519424 | 593554551711608832
	 191431067  | FleurEast       |   3174 | 2014-04-01 | 2015-04-29 | 451137254237700096 | 593529986851557377
	 615974294  | JakeQuickenden  |   3231 | 2014-06-24 | 2015-04-30 | 481374709746262016 | 593728156806316032
	 2710117976 | BlondeElectra   |   3195 | 2014-08-29 | 2015-04-30 | 505367997335470080 | 593718811376951296
	 2725900674 | LolaSaunders    |    989 | 2014-08-12 | 2015-04-30 | 499152001935351809 | 593690090049773568
	 60673341   | PaulAkister     |    598 | 2014-10-07 | 2015-04-29 | 519564953787699202 | 593390744452358144
	 87960937   | CHLOEJASMINEW   |   2365 | 2014-07-01 | 2015-04-30 | 484102432763703296 | 593703271392509952
	 2574132663 | laurenplatt7    |   1970 | 2014-05-31 | 2015-04-29 | 472798228443066369 | 593495583177965569
	 878418000  | OTYOfficial     |   1981 | 2014-11-11 | 2015-04-30 | 532157351005274112 | 593722813141364736
	 702764472  | Bhaenow         |   1582 | 2013-01-29 | 2015-04-29 | 296266615882936321 | 593532149174697984
	 2324756590 | StephanieNala   |   3235 | 2014-12-12 | 2015-04-30 | 543435863716536320 | 593702394250272769

----

	./follow_your_idol -timeline FleurEast

this scraps at most 200 tweets from the user with screen_name `@FleurEast`  
use this for the first scrap on a new user

----

	./follow_your_idol -older_or_equal_to auto -timeline FleurEast

this scraps tweets older than the ones you've cached already (another 200 tweets page at most)  
use this repetitively to fetch more past tweets from the given user

----

	./follow_your_idol -more_recent_than auto -timeline FleurEast

this scraps tweets newer than the ones you've cached already (another 200 tweets page at most)  
oftentimes one call per user suffices

----

	./follow_your_idol -fts aerosmith

does a full text search on the cached tweets for the query string `aerosmith`  
notice that the query string is expected to be a valid [tsquery expression](http://www.postgresql.org/docs/9.3/static/datatype-textsearch.html#DATATYPE-TSQUERY).

Examples:
	
	./follow_your_idol -fts "uptown+funk"
	./follow_your_idol -fts "cowell&cheryl"


## what's up with these accounts?

This is a use case I'm exploring, using the top 16 from the [X-Factor UK series 11](http://en.wikipedia.org/wiki/The_X_Factor_(UK_series_11)) contest.
