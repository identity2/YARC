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
func (a *App) InitializeAndRun(config *config.Config) {
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
	a.handler = &handlers.Handler{
		Accounts:   &models.AccountModel{DB: db},
		Comments:   &models.CommentModel{DB: db},
		Subreddits: &models.SubredditModel{DB: db},
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
	a.Get("/", a.handler.Home)
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
