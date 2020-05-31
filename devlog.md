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
