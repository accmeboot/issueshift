package main

import (
	"database/sql"
	"github.com/accmeboot/issueshift/cmd/handlers"
	"github.com/accmeboot/issueshift/cmd/helpers"
	"github.com/accmeboot/issueshift/cmd/middlewares"
	"github.com/accmeboot/issueshift/cmd/routes"
	"github.com/accmeboot/issueshift/config"
	"github.com/accmeboot/issueshift/internal/respository"
	"github.com/accmeboot/issueshift/internal/service"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.NewConfig()
	db, err := sql.Open("postgres", *cfg.DSN)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	repositoryProvider := respository.NewProvider(db)
	serviceProvider := service.NewProvider(repositoryProvider)
	helpersProvider := helpers.NewProvider()

	handlersProvider := handlers.NewProvider(serviceProvider, helpersProvider)
	middlewaresProvider := middlewares.NewProvider(serviceProvider)
	routesProvider := routes.NewProvider(handlersProvider, middlewaresProvider)

	svr := http.Server{
		Addr:         *cfg.ADDR,
		Handler:      routesProvider.Mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server at http://%s\n", *cfg.ADDR)
	log.Fatal(svr.ListenAndServe())
}
