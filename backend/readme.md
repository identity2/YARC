# Backend of YARC
This is the backend of YARC developed with Go.

## Commands
* Run project: `go run .`
* Build project: `go build .`

## API

### Authorization
Some of the requests require the user to be logged in. After logging in using `POST /login`, the user would request a token string. When performing login-required requests, the user should add the authorization header as follows:

```
Authorization: Bearer theTokenGoesHere
```

### GET `/`
Redirects to this readme page on GitHub.

### Ping

| Method | URL   | Description                | Login Required |
| ------ | ----- | -------------------------- | -------------- |
| GET    | /ping | Check if the server is up. | No             |

#### GET `/ping`
Test whether the server is up or not. If the server is up, it will reply with "ok" in the response body.

* `404 Not Found` will be returned for every failed `GET` request.

### Login & Register

| Method | URL                 | Description                | Login Required |
| ------ | ------------------- | -------------------------- | -------------- |
| POST   | /login              | Logs in a user.            | No             |
| POST   | /register           | Registers a new user.      | No             |

#### POST `/login`
Request body:

```
{
    "username": "3-20 characters",
    "password": "6-20 characters"
}
```

Response body:

If succeeded (`200`),

```
{
    "token": "the authorization token string."
}
```

Otherwise (`400`) with error message:

```
{ "error": "The reason of failure." }
```

#### POST `/register`
Request body:

```
{
    "username": "3-20 characters",
    "password": "6-20 characters",
    "email": "A valid email address"
}
```

Response body: Empty if success (`201`), otherwise (`400`) with error message.

### Article
All actions related to articles, an article can be uniquely identified by the `articleID`.

| Method | URL           | Description                | Login Required |
| ------ | ------------- | -------------------------- | -------------- |
| GET    | /article      | Return a list of articles. | No             |
| GET    | /article/{id} | Return the article {id}.   | No             |
| POST   | /article      | Insert a new article.      | Yes            |
| PUT    | /article/{id} | Update the article {id}.   | Yes            |
| DELETE | /article/{id} | Delete the article {id}.   | Yes            |

#### GET `/article`
Return a list of articles filtered by the query strings.

Request query strings:

| Query String | Value       | Default | Note                                    |
| ------------ | ----------- | ------- | --------------------------------------- |
| sort         | hot/new/old | hot     | Sort the returned list in hot/new/old.  |
| subreddit    | sub_name    |         | Return articles in subreddit.           |
| after        | articleID   |         | Return articles succeeded by articleID. |
| limit        | integer     | 25      | The number of articles to be returned.  |
| postedBy     | username    |         | Return articles posted by username.     |
| savedBy      | username    |         | Return articles saved by username.      |

Response body:

```
{
    "articles": [
        {
            "subreddit": "The subreddit the article is in.",
            "articleID": "An alphanumeric ID.",
            "type": "text/image/link",
            "body": "Text post body / Image URL / Link URL",
            "title": "The title of the article.",
            "points": 10,
            "postedBy": "username",
            "postedTime": "Date & time in string format"
        },
        ...
    ]
}
```

#### GET `/article/{id}`
Return the article {id}.

Response Body:
```
{
    "subreddit": "The subreddit the article is in.",
    "articleID": "An alphanumeric ID.",
    "type": "text/image/link",
    "body": "Text post body / Image URL / Link URL",
    "title": "The title of the article.",
    "points": 10,
    "postedBy": "username",
    "postedTime": "Date & time in string format"
}
```

#### POST `/article
Inserts a new article.

Request Body:
```
{
    "subreddit": "The subreddit the article is in.",
    "type": "text/image/link",
    "body": "Text post body / Image URL / Link URL",
    "title": "The title of the article.",
    "postedBy": "username"
}
```

Response Body:

If succeeded (`201`),

```
{"articleID": "The ID of the created article."}
```


otherwise (`401`) or (`400`) with error message.

#### PUT `/article/{id}`
Update the article {id}. Only available for text-typed posts.

Request Body:

```
{
    "body": "The updated body."
}
```

Response Body: Empty if succeeded (`200`), otherwise (`400`) or (`401`) with error message.

#### DELETE `/article/{id}`
The article {id} will be deleted.

Request Body: Empty.

Response Body: Empty if succeeded (`200`), otherwise (`400`) or (`401`) with error message.

### Search
Searches for both articles and subreddits.

#### GET `/search`
Return the search result according to the query string.

Request query strings:

| Query String | Value        |
| ------------ | ------------ |
| q            | query string |

Response Body:

```
{
    subreddits: [
        {
            "name": "The name of the subreddit.",
            "description": "The description of the subreddit."
        },
        ...
    ],
    articles: [
        {
            "subreddit": "The subreddit the article is in.",
            "articleID": "An alphanumeric ID.",
            "type": "text/image/link",
            "body": "Text post body / Image URL / Link URL",
            "title": "The title of the article.",
            "points": 10,
            "postedBy": "username",
            "postedTime": "Date & time in string format"
        },
        ...
    ]
}
```

### Comments
All actions related to comments, a comment can be uniquely identified by the `commentID`.

| Method | URL           | Description                | Login Required |
| ------ | ------------- | -------------------------- | -------------- |
| GET    | /comment      | Return a list of comments. | No             |
| GET    | /comment/{id} | Return the comment {id}.   | No             |
| POST   | /comment      | Insert a new comment.      | Yes            |
| PUT    | /comment/{id} | Update the comment {id}.   | Yes            |
| DELETE | /comment/{id} | Delete the comment {id}.   | Yes            |

#### GET `/comment`
Return a list of comments filtered by the query strings.

Request query strings:

| Query String | Value       | Default | Note                                    |
| ------------ | ----------- | ------- | --------------------------------------- |
| after        | commentID   |         | Return comments succeeded by commentID. |
| limit        | integer     | 25      | The number of comments to be returned.  |
| postedBy     | username    |         | Return comments posted by username.     |
| ofArticle    | articleID   |         | Return comments of the articleID.       |

Response Body:

```
{
    comments: [
        {
            "subreddit": "The subreddit the comment is in.",
            "articleID": "The article the comment is in.",
            "commentID": "alphanumeric ID.",
            "body": "Comment text body.",
            "points": 7,
            "postedBy": "username",
            "postedTime": "Date and time in string format."
        },
        ...
    ]
}
```

#### GET `/comment/{id}`
Return the comment {id}.

Response Body:

```
{
    "subreddit": "The subreddit the comment is in.",
    "articleID": "The article the comment is in.",
    "commentID": "alphanumeric ID.",
    "body": "Comment text body.",
    "points": 7,
    "postedBy": "username",
    "postedTime": "Date and time in string format."
}
```

#### POST `/comment`
Insert a new comment.

Request Body:

```
{
    "subreddit": "The subreddit the comment is in.",
    "articleID": "The article the comment is in.",
    "body": "Comment text body.",
    "postedBy": "username"
}
```

Response Body:

If succeeded (`201`):

```
{"commentID": "The ID of the created comment."}
```

Otherwise (`400`) or (`401`) with error message.

#### PUT `/comment/{id}`
Update the comment {id}.

Request Body:
{
    "body": "The updated text body."
}

Response Body: Empty if succeeded (`200`), otherwise (`400`) or (`401`) with error message.

#### DELETE `/comment/{id}`
Delete the comment {id}.

Request Body: Empty

Response Body: Empty if succeeded (`200`), otherwise (`400`) or (`401`) with error message.


### Account
All actions related to user accounts, an account can be uniquely identified by the `username`.

| Method | URL                  | Description                              | Login Required |
| ------ | -------------------- | ---------------------------------------- | -------------- |
| GET    | /user/{username}     | Return the info of a user.               | No             |
| PUT    | /me/bio              | Modify the bio of the current user.      | Yes            |
| POST   | /me/save/{articleID} | Save the article for the current user.   | Yes            |
| POST   | /me/join/{subreddit} | Join the subreddit for the current user. | Yes            |

#### GET `/user/{username}`
Return the user info of {username}.

Response Body:
```
{
    "username": "The username",
    "karma": 3,
    "bio": "Short bio of the username"
}
```

#### PUT `/me/bio`
Update the short bio of the current user.

Request Body:

```
{"bio": "Updated short bio of the username"}
```

Response Body: Empty if succeeded (`200`), otherwise (`401`) or (`400`) with error message.

#### POST `/me/save/{articleID}`
Save the article {articleID} for the current user.

Request Body: Empty

Response Body: Empty if succeeded (`200`), otherwise (`401`) or (`400`) with error message.

#### POST `/me/join/{subreddit}`
Join the subreddit {subreddit} for the current user.

Request Body: Empty

Response Body: Empty if succeeded (`200`), otherwise (`401`) or (`400`) with error message.

### Subreddit
All actions related to subreddits, a subreddit can be uniquely identified by its subreddit name.

| Method | URL                  | Description                              | Login Required |
| ------ | -------------------- | ---------------------------------------- | -------------- |
| GET    | /subreddit/{name}    | Return the subreddit {name}.             | No             |
| POST   | /subreddit           | Add a new subreddit.                     | Yes            |
| GET    | /trending            | Return the a list of trending subreddits.| No             |

#### GET `/subreddit/{name}`
Return the subreddit info of {name}.

Response Body:

```
{
    "name": "The name of the subreddit.",
    "description": "The description of the subreddit."
}
```

#### POST `/subreddit`
Create a new subreddit.

Request Body:

```
{
    "name": "The name of the subreddit.",
    "description": "The description of the subreddit."
}
```

Response Body: Empty if success (`201`), otherwise (`401`) or (`400`) with error message.

#### GET `/trending`
Return a list of trending subreddit.

Request query strings:

| Query String | Value       | Default | Note                                     |
| ------------ | ----------- | ------- | ---------------------------------------- |
| limit        | integer     | 5       | The number of subreddits to be returned. |

Response Body:

```
{
    subreddits: [
        {
            "name": "The name of the subreddit.",
            "description": "The description of the subreddit."
        },
        ...
    ]
}
```

### Karma
Upvote and downvote.

| Method | URL                         | Description                              | Login Required |
| ------ | --------------------------- | ---------------------------------------- | -------------- |
| POST   | /karma/article/{articleID}  | Upvote or downvote an article.           | Yes            |
| POST   | /karma/comment/{commentID}  | Upvote or downvote a comment.            | Yes            |

#### POST `/karma/article/{articleID}`
Upvote or downvote an article.

Request body:
```
{
    "articleID": "The id of the article.",
    "action": "upvote"
}
```

Response body: Empty if success (`201`), otherwise (`401`) or (`400`) with error message.

#### POST `/karma/comment/{commentID}`
Upvote or downvote a comment.

Request body:
```
{
    "commentID": "The id of the comment.",
    "action": "upvote"
}
```

Response body: Empty if success (`201`), otherwise (`401`) or (`400`) with error message.
