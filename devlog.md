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