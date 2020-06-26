# Development Log
I like to keep a chronological development log, so that I can regain my motivation after reading it.

## 2020.04.09
* Completed the header bar component.

## 2020.04.14
* Worked on the templates of `ListOfArticles.vue`, `Voter.vue`, `ArticleEntry.vue`. Almost finished, but not quite.

## 2020.04.15
* Worked on all the Right Panels, including `AboutSubreddit`, `Advertisement`, `TrendingSubreddits`.
* Set the props of `ArticleEntry`, passed from `ListOfArticles`.
* Added 10 mock articles for testing.
* Fine-tuning the layout of the main page. It's very Reddit-like now!!!

## 2020.04.16
* Completed the `CreateArticle` page. It can be linked from the `ListOfArticles` page.

## 2020.04.17
* Bug fix: Buttons on `ArticleEntry` will not trigger the article popup now.
* Completed the "expand" functionality of `ArticleEntry`.

## 2020.04.18
* Completed the `ViewArticle` page, along with the `ArticleContent`, `CommentContent` and `CreateComment` components.
* Added some mock comments for testing.

## 2020.04.19
* Now the URL query strings of `CreateArticle` will change according to the postType.
* Abandoned the overlaying approach of viewing articles. Adjusted the route paths accordingly.
* Restructured the templates: Added `Home` and `Subreddit`, while pushing `ListOfArticles` to the children.

## 2020.04.21
* Created pages: `UserProfile` and `CreateSubreddit`.
* Changed `PostingTips` to more generalized `Tips`.
* Added right panel `AboutUser`.
* Changed all `<tag></tag>` into `<tag />`.
* Changed route `/user/` to `/u/` for consistency.

## 2020.04.29
* Completed pages: `LogInRegister`, `SearchResult`.
* **Overall, the frontend templates are completed. Switching to backend development.**

## 2020.05.04
* Designed the ER Diagram and database schema.
* Started designing the backend API.

**Short Gap: Preparing for my final exams in Programming Languages, Software Engineering, Computer Networks, and Database Systems**

## 2020.05.24
* Added the edit functionality for user's short bio.
* Minor tweaks in ER Diagram and database schema (to avoid using PostgreSQL's reserved words).
* Completed `setup.sql` and `teardown.sql` according to the designed schema.

## 2020.05.25
* Added `email` and `join_time` to the `account` relations in the database.
* Designed the overall file structure of the backend server.
* Designed the RESTful API for backend. Documented in `backend/readme.md`. (Not done yet.)

## 2020.05.29
* Added an editing and deletion functionality to comments in frontend.
* Fixed a bug in frontend. Editing articles no longer pass `prop`s as `v-model`s.
* Completed the design of the RESTful API of backend. Documented in `backend/readme.md`.
* Added `UNIQUE` constraints to `aid` and `cid` in database.

## 2020.05.31
* Wrote some backend router and handler code.
* Did research on how to implement authentication.

## 2020.06.01
* Wrote some account authentication code.
* Removed to `/logout` API because I plan to use JWT, and I won't bother doing token invalidation.
* Named primary key constraints in the database.

## 2020.06.02
* Implemented middlewares: Recover Panic, Logger, Authorization.
* Completed the APIs of `/login` and `/register`.

## 2020.06.03
* Redesigned the database so that the karma system is can work correctly.
* Completed the APIs of `GET /subreddit/{name}` and `POST /subreddit`.
* Completed the APIs of `DELETE, PUT, GET /article/{id}` and `POST /article`.
* Added files mock models for writing tests.

## 2020.06.05
* Modified some HTTP response code from `200` to `204` and some from `401` to `404`.
* Wrote tests for handlers. Very tiring, but kind of learned some stuff about testing.
* Added mock data in database for testing.

## 2020.06.06
* Modified some model functions to return errors when no rows are changed under Modify and Delete operations.
* Changed `positive BOOLEAN` to `point INTEGER` in `vote_article` and `vote_comment` table of the database, so getting the points of an article could be done easier by `SUM(point)`.
* Wrote tests for models.

**Short Gap: Got addicted to Don't Starve Together, so I kept on playing and neglected the dev work. 10/10 worth it.**

## 2020.06.17
* Added `DELETE /me/save/{articleID}` and `DELETE /me/join/{subreddit}` APIs to the backend.
* Modified the implementation of `ArticleModel`'s `Get`. Using a single SQL query instead of separating the article and points queries.
* Completed the all Account related APIs. Wrote tests for all of them.
* Added more mock articles and comments so that it would be easier to test list article/comment APIs.

## 2020.06.19
* Completed writing `Get`-functions for article lists for `ArticleModel`. I struggled a bit in writing nested SQL select-queries. SQL is such a profound and complex field itself.
* Wrote tests for the newly written `Get`-functions, but not completed yet.

# 2020.06.20
* Completed writing tests for get article list functions.
* Added a comment count field in the the article response API. (This means the list article SQL I wrote yesterday was changed modified again.)
* Completed all the routes in article handler. Also wrote tests for it.
* Did some manual tests.

# 2020.06.22
* Completed the comment model and handler. Also wrote tests for them.

# 2020.06.24
* Completed the models and handlers related to Karma System. Also wrote tests for them.
* Completed the search model and handler. Also wrote tests for them.
* Added Redis connection and implemented the `/trending` subreddit API.

**Backend API is completed, now moving back into Frontend and write API-consuming code.**

# 2020.06.26
* Completed the "Page Not Found" page.
* Added the catch-all OPTIONS route to handle pre-flight CORS requests.
* Connected the backend API to the frontend login, register page.

## TODO
* Change the sorting UI of the article list in frontend to radio buttons.