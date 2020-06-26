// Package app contains the entry point of the web app and it
// defines the web app router.
package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/YuChaoGithub/YARC/backend/app/handlers"
	"github.com/YuChaoGithub/YARC/backend/app/models"
	"github.com/YuChaoGithub/YARC/backend/config"
	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver.
)

// App has a Router and a Handler instance.
// The database instance is inside the Handler instace.
type App struct {
	router  *mux.Router
	handler *handlers.Handler
}

// InitializeAndRun initializes the app with predefined configuration, and run the app.
func (a *App) InitializeAndRun(config *config.Config, jwtSecretKey string) {
	// PostgreSQL connection.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
	)
	db, err := sql.Open(config.DB.Dialect, psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Redis connection.
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.MemStore.Addr,
		Password: config.MemStore.Password,
		DB:       config.MemStore.DB,
	})
	if _, err = rdb.Ping().Result(); err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	// Inject the DBs into the models.
	a.handler = &handlers.Handler{
		Accounts:     &models.AccountModel{DB: db},
		Comments:     &models.CommentModel{DB: db},
		Subreddits:   &models.SubredditModel{DB: db, RDB: rdb},
		Articles:     &models.ArticleModel{DB: db},
		Searches:     &models.SearchModel{DB: db},
		JWTSecretKey: jwtSecretKey,
	}

	// Router.
	a.router = mux.NewRouter()
	a.setRouters()

	// Web server.
	server := &http.Server{
		Addr:         config.Server.Port,
		Handler:      a.router,
		IdleTimeout:  config.Server.IdleTimeout,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
	}

	log.Printf("Starting YARC backend server on %s\n", config.Server.Port)
	err = server.ListenAndServe()
	log.Fatal(err)
}

// setRouters sets all the routers of the API.
func (a *App) setRouters() {
	// Universal middlewares.
	a.router.Use(handlers.RecoverPanic, handlers.LogRequest, handlers.AddCORSHeader)

	// Abbreviations, to make code more succinct.
	auth := a.handler.Authorize

	// Home and ping.
	a.Get("/", a.handler.Home)
	a.Get("/ping", a.handler.Ping)

	// Authentication.
	a.Post("/login", a.handler.Login)
	a.Post("/register", a.handler.Register)

	// Article.
	a.Get("/article", a.handler.ListArticle)
	a.Get("/article/{id}", a.handler.Article)
	a.Post("/article", auth(a.handler.NewArticle))
	a.Put("/article/{id}", auth(a.handler.ModifyArticle))
	a.Delete("/article/{id}", auth(a.handler.DeleteArticle))

	// Comment.
	a.Get("/comment", a.handler.ListComment)
	a.Get("/comment/{id}", a.handler.Comment)
	a.Post("/comment", auth(a.handler.NewComment))
	a.Put("/comment/{id}", auth(a.handler.ModifyComment))
	a.Delete("/comment/{id}", auth(a.handler.DeleteComment))

	// Account.
	a.Get("/user/{username}", a.handler.User)
	a.Put("/me/bio", auth(a.handler.ModifyBio))
	a.Post("/me/save/{articleID}", auth(a.handler.SaveArticle))
	a.Delete("/me/save/{articleID}", auth(a.handler.UnsaveArticle))
	a.Post("/me/join/{subreddit}", auth(a.handler.JoinSubreddit))
	a.Delete("/me/join/{subreddit}", auth(a.handler.LeaveSubreddit))

	// Subreddit.
	a.Get("/subreddit/{name}", a.handler.Subreddit)
	a.Post("/subreddit", auth(a.handler.NewSubreddit))
	a.Get("/trending", a.handler.TrendingSubreddit)

	// Karma.
	a.Post("/karma/article/{id}", auth(a.handler.VoteArticle))
	a.Post("/karma/comment/{id}", auth(a.handler.VoteComment))

	// Search.
	a.Get("/search", a.handler.Search)

	// Options.
	a.router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}).Methods("OPTIONS")
}

// Get wraps the gorilla mux for GET method.
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the gorilla mux for POST method.
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the gorilla mux for PUT method.
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the gorilla mux for DELETE method.
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("DELETE")
}
