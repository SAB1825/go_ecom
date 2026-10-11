package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// BASE MIDDLEWARE
	r.Use(middleware.Logger)    // Logging service
	r.Use(middleware.Recoverer) // Recover from crashes

	r.Use(middleware.RequestID) // usefull for rate limiting
	r.Use(middleware.RealIP)    // use for rate limiting, analytics and tracing

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All good!!!"))
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server hash starteed at addr: %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
