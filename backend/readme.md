# Backend of YARC
This is the backend of YARC developed with Go.

## Commands
* Run project: `go run ./cmd/yarc .`
* Build project: `go build ./cmd/yarc`

## API

### Ping

| Method | URL   | Description                | Login Required |
| ------ | ----- | -------------------------- | -------------- |
| GET    | /ping | Check if the server is up. | No             |

#### GET `/ping`
Test whether the server is up or not. If the server is up, it will reply with "ok" in the response body.

* `404 Not Found` will be returned for every failed `GET` request.

### Login & Register

| Method | URL       | Description                | Login Required |
| ------ | --------- | -------------------------- | -------------- |
| POST   | /login    | Logs in a user.            | No             |
| POST   | /register | Registers a new user.      | No             |
| POST   | /logout   | Logs out a user.           | Yes            |

#### POST `/login`
Request body:

```
{
    // TODO
}
```

Response body:

Empty if success (`200`), otherwise (`400`):

```
{ "error": "The reason of failure." }
```

#### POST `/register`
Request body:

```
{
    // TODO
}
```

Response body:

Empty if success (`201`), otherwise (`400`):

```
{ "error": "The reason of failure." }
```

#### POST `/logout`
Request body: Empty

Response body:

Empty if success (`200`), otherwise (`400`):

```
{ "error": "The reason of failure." }
```

### Article
All actions related to articles, an article can be uniquely identified by the `articleID`.

| Method | URL           | Description                | Login Required |
| ------ | ------------- | -------------------------- | -------------- |
| GET    | /article      | | |
| GET    | /article/{id} | | |
| POST   | /article      | | |
| PUT    | /article/{id} | | |
| DELETE | /article/{id} | | |

#### GET `/article`
Return a list of articles filtered by the query strings.

Request query strings:

| Query String | Value       | Default | Note                                   |
| ------------ | ----------- | ------- | -------------------------------------- |
| sort         | hot/new/old | hot     |                                        |
| subreddit    | sub_name    |         |                                        |
| after        | articleID   |         | Return articles succeeded by articleID |
| limit        | integer     | 25      | The number of articles to be returned  |
| postedBy     | username    |         |                                        |
| savedBy      | username    |         |                                        |

Response body:

```
{
    // TODO
}
```

#### GET `/article/{id}`

#### POST `/article

#### PUT `/article/{id}`

#### DELETE `/article/{id}`


### Search
Searches for both articles and subreddits.

#### GET `/search`


### Comments
All actions related to articles, an article can be uniquely identified by the `articleID`.

#### GET `/comment`

#### GET `/comment/{id}`

#### POST `/comment`

#### PUT `/comment/{id}`

#### DELETE `/comment/{id}`


### Users
All actions related to users, a user can be uniquely identified by the `username`.

#### GET `/user/{id}`

