# Yet Another Reddit Clone

## What's the point?
I plan to get a web backend job after university, but I have zero web projects to show off. Therefore, I figured I might as well start building a web project to demonstrate my web skills.

## Features
* Front page article list with sorting.
* Register, login.
* Create, join subreddits.
* Post, edit, delete, vote, save articles.
* Post, edit, delete, vote comments.
* View user info: short bio, comments, articles, saved article.
* Karma system.
* Searching: subreddit and articles.

## Tech Stack
* Frontend: Vue
* Backend: Go
* Database: PostgreSQL + (maybe) Redis
* (maybe) Container: Docker

### Vue Package Used
* Vuex
* Vue-Router
* Quasar

### Go Package Used
* gorilla/mux
* lib/pq (PostgreSQL driver)
* dgrijalva/jwt-go
* teris-io/shortid,

### Why Go and Vue?
I like Go and Vue.

### Why Reddit?
I like Reddit.

## Frontend Milestones
- [x] Header bar template.
- [x] Article entry box template.
- [x] Right panel templates.
- [x] Create post page template.
- [x] View/Edit article page template.
- [x] New subreddit page template.
- [x] View profile page template.
- [x] Register/Login page template.
- [x] Search result page template.

## Backend Milestones
- [x] Design API and database schema.
- [ ] Subreddit APIs.
- [x] Article APIs.
- [x] Comment APIs.
- [ ] Karma APIs.
- [ ] Search API.
