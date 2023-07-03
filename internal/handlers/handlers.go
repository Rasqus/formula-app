package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/anras5/formula-app-backend/internal/config"
	"github.com/anras5/formula-app-backend/internal/models"
	"github.com/anras5/formula-app-backend/internal/repository"
	"github.com/anras5/formula-app-backend/internal/repository/dbrepo"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.Application
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.Application, db *sql.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Todos API is up and running",
		Version: "1.0.1",
	}

	_ = m.App.WriteJSON(w, http.StatusOK, payload)
}

// swagger:route GET /drivers/{id} driver GetRequest
// Get driver by id
//
// responses:
//
//	400: JSONResponse
//	200: Driver
func (m *Repository) OneDriver(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	driverID, err := strconv.Atoi(id)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	driver, err := m.DB.SelectDriver(driverID)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}
	_ = m.App.WriteJSON(w, http.StatusOK, driver)
}

// swagger:route POST /drivers driver PostDriverRequestBody
// Insert driver
// responses:
//
// 400: JSONResponse
// 202: PostResponse
func (m *Repository) InsertDriver(w http.ResponseWriter, r *http.Request) {
	var driver *models.Driver

	err := m.App.ReadJSON(w, r, driver)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	driver, err = m.DB.InsertDriver(*driver)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	response := config.PostResponse{
		ID:      driver.ID,
		Message: "driver inserted",
	}
	_ = m.App.WriteJSON(w, http.StatusAccepted, response)
}

// swagger:route PUT /drivers driver PutDriverRequestBody
// Update driver
// responses:
//
// 400: JSONResponse
// 202: PutResponse
func (m *Repository) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	var driver models.Driver

	err := m.App.ReadJSON(w, r, &driver)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	err = m.DB.UpdateDriver(driver)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	response := config.PutResponse{
		Message: "driver updated",
	}
	_ = m.App.WriteJSON(w, http.StatusAccepted, response)
}

// swagger:route DELETE /drivers/{id} driver DeleteRequest
// Delete driver by id
// responses:
//
// 400: JSONResponse
// 202: DeleteResponse
func (m *Repository) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	driverID, err := strconv.Atoi(id)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	err = m.DB.DeleteDriver(driverID)
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	response := config.DeleteResponse{
		Message: "driver deleted",
	}
	_ = m.App.WriteJSON(w, http.StatusAccepted, response)
}

// GraphQL is a function called when graphql request is made
func (m *Repository) GraphQL(w http.ResponseWriter, r *http.Request) {
	var req config.GraphQLRequest
	// parse request into GraphQLRequest
	err := m.App.ReadJSON(w, r, &req)

	g := NewGraph()
	g.QueryString = req.Query
	g.Variables = req.Variables

	resp, err := g.Query()
	if err != nil {
		_ = m.App.ErrorJSON(w, err)
		return
	}

	j, _ := json.MarshalIndent(resp, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(j)
}
