package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/accmeboot/issueshift/internal/data"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	queries := data.New(db)

	users, err := queries.GetAllUsers(context.Background())
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<h1>Hello</h1>"))
		if err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:         ":8000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server at :8000...")
	log.Fatal(server.ListenAndServe())
}
