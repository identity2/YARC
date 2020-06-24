# Yet Another Reddit Clone

## What's the point?
I plan to get a web backend job after university, but I have zero web projects to show off. Therefore, I figured I might as well start building a web project to demonstrate my web skills.

## Features
* Article list with new/old/hot sorting.
* Register, login by JWT.
* Create, join subreddits.
* Post, edit, delete, vote, save articles.
* Post, edit, delete, vote comments.
* View user info: short bio, comments, articles, saved article.
* Karma system.
* Searching: subreddit and articles. (Very primitive searching.)

## Tech Stack
* Frontend: Vue
* Backend: Go
* Database: PostgreSQL + Redis
* Container: Docker

### Vue Package Used
* Vuex
* Vue-Router
* Quasar

### Go Package Used
* gorilla/mux
* lib/pq (PostgreSQL driver)
* dgrijalva/jwt-go
* teris-io/shortid
* go-redis/redis

### Why Go?
I am more comfortable with static-typed languages, and I already had projects written in Java and C#, so I wanted to try out the more modern Go language. In my opinion, static-typed languages speed up the development speed, increase the readability, and reduce bugs. (change my mind)

### Why Vue?
I love the way Vue divides a document into `<template>`, `<script>`, and `<style>` sections. This improves the readability and maintainability significantly.

React code is just a bunch of mess to me, Flutter is immature for web, and I never try Angular out because it seems old.

Nevertheless, I hope WebAssembly gets more mature and popular in the future. Programming in JavaScript is...umm, a little tricky.

### Why Reddit?
I like Reddit.

(And I could demonstrate my skills on developing a medium-scaled CRUD web app by replicating it.)

### Why PostgreSQL?
Its logo is an elephant, and elephants are our friends.

Also, I think relational databases are still better than NoSQLs in general cases (especially in traditional forum sites like Reddit). It models the data in a more rigorous fashion, and makes it easier to perform complex updates and selects.

(I don't understand why people say NoSQLs "scale" better? I consider NoSQLs as useful in only very special scenarios.)

### Why Redis?
Just testing the usage of Redis. The usage of Redis in this app is minimal.

### Why Docker?
I browsed the job ads in my local area and it seemed to be a good plus to know Docker.

# Planned Completion Time
Mid-to-Late-July, 2020