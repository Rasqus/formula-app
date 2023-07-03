package main

import (
	"github.com/anras5/formula-app-backend/internal/config"
	"github.com/anras5/formula-app-backend/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	apiMid "github.com/go-openapi/runtime/middleware"
	"net/http"
)

func routes(app *config.Application) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	//mux.Use(EnableCORS)

	// REST
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/drivers/{id}", handlers.Repo.OneDriver)
	mux.Post("/drivers", handlers.Repo.InsertDriver)
	mux.Put("/drivers", handlers.Repo.UpdateDriver)
	mux.Delete("/drivers/{id}", handlers.Repo.DeleteDriver)

	// GRAPHQL
	mux.Post("/graphql", handlers.Repo.GraphQL)

	// Swagger documentation
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	opts := apiMid.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	sh := apiMid.SwaggerUI(opts, nil)

	mux.Get("/docs", func(writer http.ResponseWriter, request *http.Request) {
		sh.ServeHTTP(writer, request)
	})

	return mux
}
