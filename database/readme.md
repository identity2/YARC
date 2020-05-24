# Database
This directory contains the SQL scripts for the backend database.

# Design
## ER Diagram
![ER Diagram](https://i.imgur.com/hU1nP2a.jpg)

## Schema
(**key: bold**, *foreign key: italic*)

* account(**username**, hashed_password, karma, bio)
* subreddit(**sub_name**, description)
* join_sub(***username***, ***sub_name***)
* article(***sub_name***, **aid**, type, body, title, points, *posted_by*, posted_time)
* comment(***sub_name***, ***aid***, **cid**, body, points, *posted_by*, posted_time)
* save_article(***username***, ***sub_name***, ***aid***)
