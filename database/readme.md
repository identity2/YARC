# Database
This directory contains the SQL scripts for the backend database.

# Design
## ER Diagram
![ER Diagram](https://i.imgur.com/ohpvfpO.jpg)

## Schema
(**key: bold**, *foreign key: italic*)

* account(**username**, hashed_password, email, bio, join_time)
* subreddit(**sub_name**, description)
* join_sub(***username***, ***sub_name***)
* article(***sub_name***, **aid**, type, body, title, *posted_by*, posted_time)
* comment(***sub_name***, ***aid***, **cid**, body, *posted_by*, posted_time)
* save_article(***username***, ***sub_name***, ***aid***)
* vote_article(***username***, ***sub_name***, ***aid***, point)
* vote_comment(***username***, ***sub_name***, ***aid***, ***cid***, point)

## Notes
* No indexes are created because the data size is rather small.
* Redis is used for tracking trending subreddits only because my purpose is to just try it out.

## Retrospect
* Maybe I could merge the article entity together with the comment entity for simplicity. Also, vote_article and vote_comment could definitely be merged.
* I could use Redis to cache hot articles and karma, but I was too lazy to do it.
