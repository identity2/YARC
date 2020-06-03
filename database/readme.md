# Database
This directory contains the SQL scripts for the backend database.

# Design
## ER Diagram
![ER Diagram](https://i.imgur.com/XXgVUqa.jpg)

## Schema
(**key: bold**, *foreign key: italic*)

* account(**username**, hashed_password, email, bio, join_time)
* subreddit(**sub_name**, description)
* join_sub(***username***, ***sub_name***)
* article(***sub_name***, **aid**, type, body, title, *posted_by*, posted_time)
* comment(***sub_name***, ***aid***, **cid**, body, *posted_by*, posted_time)
* save_article(***username***, ***sub_name***, ***aid***)
* vote_article(***username***, ***sub_name***, ***aid***, positive)
* vote_comment(***username***, ***sub_name***, ***aid***, ***cid***, positive)
