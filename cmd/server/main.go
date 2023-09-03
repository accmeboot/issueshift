package main

import (
	"github.com/accmeboot/issueshift/internal/config"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.NewConfig()
	db, err := GetDatabase(*cfg.DSN)
	if err != nil {
		panic(err)
	}

	router := NewRouter(db)

	router.MapRoutes()

	server := http.Server{
		Addr:         *cfg.ADDR,
		Handler:      router.Mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server at http://%s\n", *cfg.ADDR)
	log.Fatal(server.ListenAndServe())
}
