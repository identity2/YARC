![Backend Build and Test](https://github.com/YuChaoGithub/YARC/workflows/Backend%20Build%20and%20Test/badge.svg)

# Yet Another Reddit Clone

[Website](https://yarc-29bed.firebaseapp.com/) ． [Demo Video](https://youtu.be/BwcGhgvhw6k) ． [Blog Post](https://shinerightstudio.com/posts/project-yarc-cloning-reddit-using-go-and-vue/)

![Homepage](https://i.imgur.com/jvVG5yU.png)

## What's the point?
I plan to get a web backend job after university, but I have zero web projects to show off. Therefore, I figured I might as well start building a web project to demonstrate my web skills.

## Features
* Article list with new/old/hot sorting.
* Register, login by JWT.
* Create, join subreddits.
* Post text, image, link articles. (Using Imgur API for image posts.)
* Edit, delete, vote, save articles.
* Post, edit, delete, vote comments.
* View user info: short bio, comments, articles, saved article.
* Karma system.
* Searching: subreddit and articles. (Very primitive searching.)
* [Backend API](https://github.com/YuChaoGithub/YARC/blob/master/backend/readme.md)

## Tech Stack
* Frontend: Vue
* Backend: Go
* Database: PostgreSQL + Redis
* Container: Docker
* Deployed on: Heroku (backend & database) + Firebase (frontend)

## Run the project
Make sure the `docker` command is available in your environment.

### Start Development Server
* Backend server, database, and Redis: `docker-compose up -d` (Run `docker-compose down` to clean up afterwards.)
* Frontend server: `cd frontend` and `npm run serve` (It would probably run on port `8081` since `8080` is occupied by the backend.)

*(There seems to be a problem with the quasar dependency when running on Docker, making it impossible to run a hot-reload frontend server in the container. ¯\\_(ツ)_/¯)*

### Run Backend Tests
```
cd backend
chmod +x test.sh clean_test.sh
./test.sh && ./clean_test.sh
cd ..
```

### Vue & npm Package Used
* Vuex
* Vue-Router
* Quasar
* Axios
* Moment

### Go Package Used
* gorilla/mux
* lib/pq (PostgreSQL driver)
* dgrijalva/jwt-go
* teris-io/shortid
* go-redis/redis

### Why Go?
I am more comfortable with static-typed languages, and I already had projects written in Java and C#, so I wanted to try out the more modern Go language. In my opinion, static-typed languages are less error-prone, more readable, and thus they speed up the development process. (change my mind)

### Why Vue?
I love the way Vue divides a document into `<template>`, `<script>`, `<style>` sections, and the properties of a Vue object make developers organize code in a clear way. This improves the readability and maintainability significantly.

React code is just a bunch of mess to me, Flutter is immature for web, and I never try Angular out because it seems old.

Nevertheless, I hope WebAssembly gets more mature and popular in the future. Programming in JavaScript is...umm, a little tricky.

### Why Reddit?
I like Reddit.

(And I could demonstrate my skills on developing a medium-scaled CRUD web app by replicating it.)

### Why PostgreSQL?
Most relational databases work pretty much the same except in minor details. I went with PostgreSQL because it was arguably the most popular relational database these couple years.

Why relational databases over NoSQL though? I argue that people should always stick to relational database as the main storage for their web app, only in very special cases should NoSQL come into the picture and *complements* the main relational database. Almost all data are relational, so relational databases simply model real world data better and more rigorously. It also makes it much simpler to perform complex updates and selects. All in all, I don't agree with the point of view that NoSQL "scales" better. To me, NoSQL is just a complementary tool for storing documents.

### Why Redis?
Just testing the usage of Redis. The usage of Redis in this app is minimal. Could definite explore Redis more on future projects.

### Why Docker?
It takes some time to make it work correctly, but saves more time in the long run. Also, I browsed the job ads in my local area and it seemed to be a good plus to know Docker.

## Screenshots
### Search
![Search](https://i.imgur.com/H21JMSO.png)

### Posting
![Posting](https://i.imgur.com/my7QdxU.png)

### Text Article
![Article Page (text article)](https://i.imgur.com/nYRDvol.png)

### Link Article
![Article Page (link article)](https://i.imgur.com/Fy0y0Dz.png)

### Article List
![Article List](https://i.imgur.com/018xBow.png)

## In Retrospect...
### Database
* The article entity and the comment entity could be merged into a single "post" entity with a "parent" field. This makes it easier to add new similar entities to the system and makes nested posts possible. Also, I wouldn't need to separate the vote_article and vote_comment into different tables by applying this approach.

### Backend
* There are still some N+1 problems using this API, for example, when the frontend fetches the article list of N articles, it then has to send N requests separately to fetch the voting state and the join state of each article. I could have added the vote state as well as the join state together with the article API.
* I could use Redis to cache hot articles.

### Frontend
* The naming of the components could be better.
* Many components could be merged and reused. I didn't plan too good at the beginner.

(This is just a showcase project so I would not refactor the code. However, I will make sure I don't make the same design mistakes in future projects.)