package main

import (
	"fmt"
	"github.com/anras5/formula-app-backend/internal/config"
	"github.com/anras5/formula-app-backend/internal/driver"
	"github.com/anras5/formula-app-backend/internal/handlers"
	"log"
	"net/http"
	"os"
)

const port = ":8080"

var app config.Application

func main() {

	// -------------------------------------------------------------------------------------------- //
	// Connect to database
	log.Println("Connecting to database on port 5432...")
	db, err := driver.ConnectSQL(fmt.Sprintf("host=postgres-db port=5432 dbname=formuladb user=%s password=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("SSL_MODE")))
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()

	// -------------------------------------------------------------------------------------------- //
	// set repo and handlers
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
